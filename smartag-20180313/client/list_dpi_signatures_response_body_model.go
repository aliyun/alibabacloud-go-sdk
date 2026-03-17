// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDpiSignaturesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDpiSignature(v []*ListDpiSignaturesResponseBodyDpiSignature) *ListDpiSignaturesResponseBody
	GetDpiSignature() []*ListDpiSignaturesResponseBodyDpiSignature
	SetNextToken(v string) *ListDpiSignaturesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDpiSignaturesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDpiSignaturesResponseBody
	GetTotalCount() *int32
}

type ListDpiSignaturesResponseBody struct {
	// The information about the application.
	DpiSignature []*ListDpiSignaturesResponseBodyDpiSignature `json:"DpiSignature,omitempty" xml:"DpiSignature,omitempty" type:"Repeated"`
	// The token returned for the next query.
	//
	// example:
	//
	// FFrMV38kR4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 63081123-B7C0-4BC9-B9E5-59E77A616EC9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned on the current page.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDpiSignaturesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDpiSignaturesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDpiSignaturesResponseBody) GetDpiSignature() []*ListDpiSignaturesResponseBodyDpiSignature {
	return s.DpiSignature
}

func (s *ListDpiSignaturesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDpiSignaturesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDpiSignaturesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDpiSignaturesResponseBody) SetDpiSignature(v []*ListDpiSignaturesResponseBodyDpiSignature) *ListDpiSignaturesResponseBody {
	s.DpiSignature = v
	return s
}

func (s *ListDpiSignaturesResponseBody) SetNextToken(v string) *ListDpiSignaturesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDpiSignaturesResponseBody) SetRequestId(v string) *ListDpiSignaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDpiSignaturesResponseBody) SetTotalCount(v int32) *ListDpiSignaturesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDpiSignaturesResponseBody) Validate() error {
	if s.DpiSignature != nil {
		for _, item := range s.DpiSignature {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDpiSignaturesResponseBodyDpiSignature struct {
	// The ID of the application group to which the application belongs.
	//
	// example:
	//
	// 20
	DpiGroupId *string `json:"DpiGroupId,omitempty" xml:"DpiGroupId,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// 235
	DpiSignatureId *string `json:"DpiSignatureId,omitempty" xml:"DpiSignatureId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// EdgeCast
	DpiSignatureName *string `json:"DpiSignatureName,omitempty" xml:"DpiSignatureName,omitempty"`
	// The earliest version of engine that supports the application.
	//
	// example:
	//
	// 0-0.0.1
	MinEngineVersion *string `json:"MinEngineVersion,omitempty" xml:"MinEngineVersion,omitempty"`
	// The earliest version of signature database that supports the application.
	//
	// example:
	//
	// 20201117_1_0-0.0.1
	MinSignatureDbVersion *string `json:"MinSignatureDbVersion,omitempty" xml:"MinSignatureDbVersion,omitempty"`
}

func (s ListDpiSignaturesResponseBodyDpiSignature) String() string {
	return dara.Prettify(s)
}

func (s ListDpiSignaturesResponseBodyDpiSignature) GoString() string {
	return s.String()
}

func (s *ListDpiSignaturesResponseBodyDpiSignature) GetDpiGroupId() *string {
	return s.DpiGroupId
}

func (s *ListDpiSignaturesResponseBodyDpiSignature) GetDpiSignatureId() *string {
	return s.DpiSignatureId
}

func (s *ListDpiSignaturesResponseBodyDpiSignature) GetDpiSignatureName() *string {
	return s.DpiSignatureName
}

func (s *ListDpiSignaturesResponseBodyDpiSignature) GetMinEngineVersion() *string {
	return s.MinEngineVersion
}

func (s *ListDpiSignaturesResponseBodyDpiSignature) GetMinSignatureDbVersion() *string {
	return s.MinSignatureDbVersion
}

func (s *ListDpiSignaturesResponseBodyDpiSignature) SetDpiGroupId(v string) *ListDpiSignaturesResponseBodyDpiSignature {
	s.DpiGroupId = &v
	return s
}

func (s *ListDpiSignaturesResponseBodyDpiSignature) SetDpiSignatureId(v string) *ListDpiSignaturesResponseBodyDpiSignature {
	s.DpiSignatureId = &v
	return s
}

func (s *ListDpiSignaturesResponseBodyDpiSignature) SetDpiSignatureName(v string) *ListDpiSignaturesResponseBodyDpiSignature {
	s.DpiSignatureName = &v
	return s
}

func (s *ListDpiSignaturesResponseBodyDpiSignature) SetMinEngineVersion(v string) *ListDpiSignaturesResponseBodyDpiSignature {
	s.MinEngineVersion = &v
	return s
}

func (s *ListDpiSignaturesResponseBodyDpiSignature) SetMinSignatureDbVersion(v string) *ListDpiSignaturesResponseBodyDpiSignature {
	s.MinSignatureDbVersion = &v
	return s
}

func (s *ListDpiSignaturesResponseBodyDpiSignature) Validate() error {
	return dara.Validate(s)
}
