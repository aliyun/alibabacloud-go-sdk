// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadAICPublicKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UploadAICPublicKeyResponseBody
	GetRequestId() *string
}

type UploadAICPublicKeyResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 125B04C7-3D0D-4245-AF96-14E3758E3F06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadAICPublicKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadAICPublicKeyResponseBody) GoString() string {
	return s.String()
}

func (s *UploadAICPublicKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadAICPublicKeyResponseBody) SetRequestId(v string) *UploadAICPublicKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadAICPublicKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
