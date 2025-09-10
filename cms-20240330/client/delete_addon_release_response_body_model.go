// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAddonReleaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAddonReleaseResponseBody
	GetRequestId() *string
}

type DeleteAddonReleaseResponseBody struct {
	// example:
	//
	// 264C3E89-BE6E-5F82-A484-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteAddonReleaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAddonReleaseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAddonReleaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAddonReleaseResponseBody) SetRequestId(v string) *DeleteAddonReleaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAddonReleaseResponseBody) Validate() error {
	return dara.Validate(s)
}
