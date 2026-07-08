// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVirtualThreeElementsVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *VirtualThreeElementsVerificationResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *VirtualThreeElementsVerificationResponseBody
	GetCode() *string
	SetData(v *VirtualThreeElementsVerificationResponseBodyData) *VirtualThreeElementsVerificationResponseBody
	GetData() *VirtualThreeElementsVerificationResponseBodyData
	SetMessage(v string) *VirtualThreeElementsVerificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *VirtualThreeElementsVerificationResponseBody
	GetRequestId() *string
}

type VirtualThreeElementsVerificationResponseBody struct {
	AccessDeniedDetail *string                                           `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *VirtualThreeElementsVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message            *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VirtualThreeElementsVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VirtualThreeElementsVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *VirtualThreeElementsVerificationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *VirtualThreeElementsVerificationResponseBody) GetCode() *string {
	return s.Code
}

func (s *VirtualThreeElementsVerificationResponseBody) GetData() *VirtualThreeElementsVerificationResponseBodyData {
	return s.Data
}

func (s *VirtualThreeElementsVerificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VirtualThreeElementsVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VirtualThreeElementsVerificationResponseBody) SetAccessDeniedDetail(v string) *VirtualThreeElementsVerificationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *VirtualThreeElementsVerificationResponseBody) SetCode(v string) *VirtualThreeElementsVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *VirtualThreeElementsVerificationResponseBody) SetData(v *VirtualThreeElementsVerificationResponseBodyData) *VirtualThreeElementsVerificationResponseBody {
	s.Data = v
	return s
}

func (s *VirtualThreeElementsVerificationResponseBody) SetMessage(v string) *VirtualThreeElementsVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *VirtualThreeElementsVerificationResponseBody) SetRequestId(v string) *VirtualThreeElementsVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *VirtualThreeElementsVerificationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VirtualThreeElementsVerificationResponseBodyData struct {
	// example:
	//
	// 81
	IsConsistent *int64 `json:"IsConsistent,omitempty" xml:"IsConsistent,omitempty"`
}

func (s VirtualThreeElementsVerificationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s VirtualThreeElementsVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *VirtualThreeElementsVerificationResponseBodyData) GetIsConsistent() *int64 {
	return s.IsConsistent
}

func (s *VirtualThreeElementsVerificationResponseBodyData) SetIsConsistent(v int64) *VirtualThreeElementsVerificationResponseBodyData {
	s.IsConsistent = &v
	return s
}

func (s *VirtualThreeElementsVerificationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
