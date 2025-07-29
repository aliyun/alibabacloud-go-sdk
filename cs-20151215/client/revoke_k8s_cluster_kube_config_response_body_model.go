// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeK8sClusterKubeConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
}

type RevokeK8sClusterKubeConfigResponseBody struct {
}

func (s RevokeK8sClusterKubeConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeK8sClusterKubeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeK8sClusterKubeConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
