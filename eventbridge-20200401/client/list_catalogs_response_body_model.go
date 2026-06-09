// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCatalogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCatalogsResponseBody
	GetCode() *string
	SetData(v *ListCatalogsResponseBodyData) *ListCatalogsResponseBody
	GetData() *ListCatalogsResponseBodyData
	SetMessage(v string) *ListCatalogsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCatalogsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCatalogsResponseBody
	GetSuccess() *bool
}

type ListCatalogsResponseBody struct {
	// example:
	//
	// 200
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListCatalogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCatalogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCatalogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCatalogsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCatalogsResponseBody) GetData() *ListCatalogsResponseBodyData {
	return s.Data
}

func (s *ListCatalogsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCatalogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCatalogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCatalogsResponseBody) SetCode(v string) *ListCatalogsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCatalogsResponseBody) SetData(v *ListCatalogsResponseBodyData) *ListCatalogsResponseBody {
	s.Data = v
	return s
}

func (s *ListCatalogsResponseBody) SetMessage(v string) *ListCatalogsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCatalogsResponseBody) SetRequestId(v string) *ListCatalogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCatalogsResponseBody) SetSuccess(v bool) *ListCatalogsResponseBody {
	s.Success = &v
	return s
}

func (s *ListCatalogsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCatalogsResponseBodyData struct {
	Catalogs []*Catalog `json:"Catalogs,omitempty" xml:"Catalogs,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListCatalogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCatalogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCatalogsResponseBodyData) GetCatalogs() []*Catalog {
	return s.Catalogs
}

func (s *ListCatalogsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCatalogsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListCatalogsResponseBodyData) SetCatalogs(v []*Catalog) *ListCatalogsResponseBodyData {
	s.Catalogs = v
	return s
}

func (s *ListCatalogsResponseBodyData) SetNextToken(v string) *ListCatalogsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListCatalogsResponseBodyData) SetTotal(v int32) *ListCatalogsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListCatalogsResponseBodyData) Validate() error {
	if s.Catalogs != nil {
		for _, item := range s.Catalogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
