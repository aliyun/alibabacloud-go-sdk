// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadPCACertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifier(v string) *UploadPCACertResponseBody
	GetIdentifier() *string
	SetRequestId(v string) *UploadPCACertResponseBody
	GetRequestId() *string
}

type UploadPCACertResponseBody struct {
	// The certificate identifier.
	//
	// example:
	//
	// 1ed65580-7e33-6a50-8630-dd13fdc009ee
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadPCACertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadPCACertResponseBody) GoString() string {
	return s.String()
}

func (s *UploadPCACertResponseBody) GetIdentifier() *string {
	return s.Identifier
}

func (s *UploadPCACertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadPCACertResponseBody) SetIdentifier(v string) *UploadPCACertResponseBody {
	s.Identifier = &v
	return s
}

func (s *UploadPCACertResponseBody) SetRequestId(v string) *UploadPCACertResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadPCACertResponseBody) Validate() error {
	return dara.Validate(s)
}
