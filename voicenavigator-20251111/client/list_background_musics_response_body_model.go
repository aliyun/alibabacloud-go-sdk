// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBackgroundMusicsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBackgroundMusicsResponseBody
	GetCode() *string
	SetData(v *ListBackgroundMusicsResponseBodyData) *ListBackgroundMusicsResponseBody
	GetData() *ListBackgroundMusicsResponseBodyData
	SetHttpStatusCode(v int32) *ListBackgroundMusicsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListBackgroundMusicsResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListBackgroundMusicsResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListBackgroundMusicsResponseBody
	GetRequestId() *string
}

type ListBackgroundMusicsResponseBody struct {
	// example:
	//
	// OK
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListBackgroundMusicsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance af81a389-91f0-4157-8d82-720edd02b66a
	//
	//  does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 7401D698-0AAC-5909-B68E-88C68805FFB8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBackgroundMusicsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBackgroundMusicsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBackgroundMusicsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBackgroundMusicsResponseBody) GetData() *ListBackgroundMusicsResponseBodyData {
	return s.Data
}

func (s *ListBackgroundMusicsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBackgroundMusicsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBackgroundMusicsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListBackgroundMusicsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBackgroundMusicsResponseBody) SetCode(v string) *ListBackgroundMusicsResponseBody {
	s.Code = &v
	return s
}

func (s *ListBackgroundMusicsResponseBody) SetData(v *ListBackgroundMusicsResponseBodyData) *ListBackgroundMusicsResponseBody {
	s.Data = v
	return s
}

func (s *ListBackgroundMusicsResponseBody) SetHttpStatusCode(v int32) *ListBackgroundMusicsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBackgroundMusicsResponseBody) SetMessage(v string) *ListBackgroundMusicsResponseBody {
	s.Message = &v
	return s
}

func (s *ListBackgroundMusicsResponseBody) SetParams(v []*string) *ListBackgroundMusicsResponseBody {
	s.Params = v
	return s
}

func (s *ListBackgroundMusicsResponseBody) SetRequestId(v string) *ListBackgroundMusicsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBackgroundMusicsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListBackgroundMusicsResponseBodyData struct {
	BackgroundMusics []*ListBackgroundMusicsResponseBodyDataBackgroundMusics `json:"BackgroundMusics,omitempty" xml:"BackgroundMusics,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBackgroundMusicsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBackgroundMusicsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBackgroundMusicsResponseBodyData) GetBackgroundMusics() []*ListBackgroundMusicsResponseBodyDataBackgroundMusics {
	return s.BackgroundMusics
}

func (s *ListBackgroundMusicsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBackgroundMusicsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBackgroundMusicsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBackgroundMusicsResponseBodyData) SetBackgroundMusics(v []*ListBackgroundMusicsResponseBodyDataBackgroundMusics) *ListBackgroundMusicsResponseBodyData {
	s.BackgroundMusics = v
	return s
}

func (s *ListBackgroundMusicsResponseBodyData) SetPageNumber(v int32) *ListBackgroundMusicsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListBackgroundMusicsResponseBodyData) SetPageSize(v int32) *ListBackgroundMusicsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListBackgroundMusicsResponseBodyData) SetTotalCount(v int32) *ListBackgroundMusicsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListBackgroundMusicsResponseBodyData) Validate() error {
	if s.BackgroundMusics != nil {
		for _, item := range s.BackgroundMusics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBackgroundMusicsResponseBodyDataBackgroundMusics struct {
	// ID
	//
	// example:
	//
	// 3258b551-4847-45fa-bbd8-838d90b90080
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListBackgroundMusicsResponseBodyDataBackgroundMusics) String() string {
	return dara.Prettify(s)
}

func (s ListBackgroundMusicsResponseBodyDataBackgroundMusics) GoString() string {
	return s.String()
}

func (s *ListBackgroundMusicsResponseBodyDataBackgroundMusics) GetId() *string {
	return s.Id
}

func (s *ListBackgroundMusicsResponseBodyDataBackgroundMusics) GetName() *string {
	return s.Name
}

func (s *ListBackgroundMusicsResponseBodyDataBackgroundMusics) SetId(v string) *ListBackgroundMusicsResponseBodyDataBackgroundMusics {
	s.Id = &v
	return s
}

func (s *ListBackgroundMusicsResponseBodyDataBackgroundMusics) SetName(v string) *ListBackgroundMusicsResponseBodyDataBackgroundMusics {
	s.Name = &v
	return s
}

func (s *ListBackgroundMusicsResponseBodyDataBackgroundMusics) Validate() error {
	return dara.Validate(s)
}
