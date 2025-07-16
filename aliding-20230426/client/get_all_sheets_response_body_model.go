// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAllSheetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAllSheetsResponseBody
	GetRequestId() *string
	SetValue(v []*GetAllSheetsResponseBodyValue) *GetAllSheetsResponseBody
	GetValue() []*GetAllSheetsResponseBodyValue
}

type GetAllSheetsResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// []
	Value []*GetAllSheetsResponseBodyValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s GetAllSheetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAllSheetsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllSheetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAllSheetsResponseBody) GetValue() []*GetAllSheetsResponseBodyValue {
	return s.Value
}

func (s *GetAllSheetsResponseBody) SetRequestId(v string) *GetAllSheetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAllSheetsResponseBody) SetValue(v []*GetAllSheetsResponseBodyValue) *GetAllSheetsResponseBody {
	s.Value = v
	return s
}

func (s *GetAllSheetsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAllSheetsResponseBodyValue struct {
	// example:
	//
	// stxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// Sheet1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetAllSheetsResponseBodyValue) String() string {
	return dara.Prettify(s)
}

func (s GetAllSheetsResponseBodyValue) GoString() string {
	return s.String()
}

func (s *GetAllSheetsResponseBodyValue) GetId() *string {
	return s.Id
}

func (s *GetAllSheetsResponseBodyValue) GetName() *string {
	return s.Name
}

func (s *GetAllSheetsResponseBodyValue) SetId(v string) *GetAllSheetsResponseBodyValue {
	s.Id = &v
	return s
}

func (s *GetAllSheetsResponseBodyValue) SetName(v string) *GetAllSheetsResponseBodyValue {
	s.Name = &v
	return s
}

func (s *GetAllSheetsResponseBodyValue) Validate() error {
	return dara.Validate(s)
}
