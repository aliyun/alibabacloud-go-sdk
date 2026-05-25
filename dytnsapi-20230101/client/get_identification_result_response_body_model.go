// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentificationResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetIdentificationResultResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetIdentificationResultResponseBody
	GetCode() *string
	SetData(v *GetIdentificationResultResponseBodyData) *GetIdentificationResultResponseBody
	GetData() *GetIdentificationResultResponseBodyData
	SetMessage(v string) *GetIdentificationResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetIdentificationResultResponseBody
	GetRequestId() *string
}

type GetIdentificationResultResponseBody struct {
	// example:
	//
	// -
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetIdentificationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 68A40250-50CD-034C-B728-0BD135850177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIdentificationResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIdentificationResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetIdentificationResultResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetIdentificationResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetIdentificationResultResponseBody) GetData() *GetIdentificationResultResponseBodyData {
	return s.Data
}

func (s *GetIdentificationResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetIdentificationResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIdentificationResultResponseBody) SetAccessDeniedDetail(v string) *GetIdentificationResultResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetIdentificationResultResponseBody) SetCode(v string) *GetIdentificationResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetIdentificationResultResponseBody) SetData(v *GetIdentificationResultResponseBodyData) *GetIdentificationResultResponseBody {
	s.Data = v
	return s
}

func (s *GetIdentificationResultResponseBody) SetMessage(v string) *GetIdentificationResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetIdentificationResultResponseBody) SetRequestId(v string) *GetIdentificationResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIdentificationResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIdentificationResultResponseBodyData struct {
	// example:
	//
	// true
	IsIdentified *bool `json:"IsIdentified,omitempty" xml:"IsIdentified,omitempty"`
}

func (s GetIdentificationResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetIdentificationResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIdentificationResultResponseBodyData) GetIsIdentified() *bool {
	return s.IsIdentified
}

func (s *GetIdentificationResultResponseBodyData) SetIsIdentified(v bool) *GetIdentificationResultResponseBodyData {
	s.IsIdentified = &v
	return s
}

func (s *GetIdentificationResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
