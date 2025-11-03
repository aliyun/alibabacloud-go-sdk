// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetResourcesResponseBody
	GetCode() *string
	SetData(v *GetResourcesResponseBodyData) *GetResourcesResponseBody
	GetData() *GetResourcesResponseBodyData
	SetMessage(v string) *GetResourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetResourcesResponseBody
	GetRequestId() *string
}

type GetResourcesResponseBody struct {
	// example:
	//
	// Success
	Code *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetResourcesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// result: code=1 msg=(Request failed, status_code != 200)
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 35F91AAB-5FDF-5A22-B211-C7C6B00817D0
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s GetResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetResourcesResponseBody) GetData() *GetResourcesResponseBodyData {
	return s.Data
}

func (s *GetResourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourcesResponseBody) SetCode(v string) *GetResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *GetResourcesResponseBody) SetData(v *GetResourcesResponseBodyData) *GetResourcesResponseBody {
	s.Data = v
	return s
}

func (s *GetResourcesResponseBody) SetMessage(v string) *GetResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *GetResourcesResponseBody) SetRequestId(v string) *GetResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourcesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourcesResponseBodyData struct {
	// example:
	//
	// 2354
	Total *float32 `json:"total,omitempty" xml:"total,omitempty"`
	// example:
	//
	// Kbytes
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// example:
	//
	// 100
	Usage *float32 `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s GetResourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResourcesResponseBodyData) GetTotal() *float32 {
	return s.Total
}

func (s *GetResourcesResponseBodyData) GetUnit() *string {
	return s.Unit
}

func (s *GetResourcesResponseBodyData) GetUsage() *float32 {
	return s.Usage
}

func (s *GetResourcesResponseBodyData) SetTotal(v float32) *GetResourcesResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetResourcesResponseBodyData) SetUnit(v string) *GetResourcesResponseBodyData {
	s.Unit = &v
	return s
}

func (s *GetResourcesResponseBodyData) SetUsage(v float32) *GetResourcesResponseBodyData {
	s.Usage = &v
	return s
}

func (s *GetResourcesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
