// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSignatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSignatureResponseBody
	GetRequestId() *string
	SetSignatureId(v string) *CreateSignatureResponseBody
	GetSignatureId() *string
	SetSignatureName(v string) *CreateSignatureResponseBody
	GetSignatureName() *string
}

type CreateSignatureResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the back-end signature key.
	//
	// example:
	//
	// dd05f1c54d6749eda95f9fa6d491449a
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// The name of the back-end signature key.
	//
	// example:
	//
	// backendsignature
	SignatureName *string `json:"SignatureName,omitempty" xml:"SignatureName,omitempty"`
}

func (s CreateSignatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSignatureResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSignatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSignatureResponseBody) GetSignatureId() *string {
	return s.SignatureId
}

func (s *CreateSignatureResponseBody) GetSignatureName() *string {
	return s.SignatureName
}

func (s *CreateSignatureResponseBody) SetRequestId(v string) *CreateSignatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSignatureResponseBody) SetSignatureId(v string) *CreateSignatureResponseBody {
	s.SignatureId = &v
	return s
}

func (s *CreateSignatureResponseBody) SetSignatureName(v string) *CreateSignatureResponseBody {
	s.SignatureName = &v
	return s
}

func (s *CreateSignatureResponseBody) Validate() error {
	return dara.Validate(s)
}
