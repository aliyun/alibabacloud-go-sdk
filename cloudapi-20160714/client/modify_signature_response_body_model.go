// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySignatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySignatureResponseBody
	GetRequestId() *string
	SetSignatureId(v string) *ModifySignatureResponseBody
	GetSignatureId() *string
	SetSignatureName(v string) *ModifySignatureResponseBody
	GetSignatureName() *string
}

type ModifySignatureResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the key.
	//
	// example:
	//
	// dd05f1c54d6749eda95f9fa6d491449a
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// The name of the key.
	//
	// example:
	//
	// backendsignature
	SignatureName *string `json:"SignatureName,omitempty" xml:"SignatureName,omitempty"`
}

func (s ModifySignatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySignatureResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySignatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySignatureResponseBody) GetSignatureId() *string {
	return s.SignatureId
}

func (s *ModifySignatureResponseBody) GetSignatureName() *string {
	return s.SignatureName
}

func (s *ModifySignatureResponseBody) SetRequestId(v string) *ModifySignatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySignatureResponseBody) SetSignatureId(v string) *ModifySignatureResponseBody {
	s.SignatureId = &v
	return s
}

func (s *ModifySignatureResponseBody) SetSignatureName(v string) *ModifySignatureResponseBody {
	s.SignatureName = &v
	return s
}

func (s *ModifySignatureResponseBody) Validate() error {
	return dara.Validate(s)
}
