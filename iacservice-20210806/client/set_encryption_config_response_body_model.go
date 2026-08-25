// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetEncryptionConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetEncryptionConfigResponseBody
	GetRequestId() *string
}

type SetEncryptionConfigResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 7FA0FF4A-ABD4-54F6-BEAC-B4273EBA10A2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SetEncryptionConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetEncryptionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetEncryptionConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetEncryptionConfigResponseBody) SetRequestId(v string) *SetEncryptionConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetEncryptionConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
