package olm

import (
	"crypto/ed25519"
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUtilEd25519PKToCurve25519(t *testing.T) {
	pub, _, err := ed25519.GenerateKey(rand.Reader)
	require.Nil(t, err)

	_, err = Ed25519PKToCurve25519(pub)
	require.Nil(t, err)
}

func TestUtilEd25519SKToCurve25519(t *testing.T) {
	_, priv, err := ed25519.GenerateKey(rand.Reader)
	require.Nil(t, err)

	_, err = Ed25519SKToCurve25519(priv)
	require.Nil(t, err)
}
