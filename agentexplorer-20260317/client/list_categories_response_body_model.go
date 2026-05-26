// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCategoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Category) *ListCategoriesResponseBody
	GetData() []*Category
	SetMessage(v string) *ListCategoriesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCategoriesResponseBody
	GetRequestId() *string
}

type ListCategoriesResponseBody struct {
	Data []*Category `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 099A671E-FA21-5A36-8A73-918572DDEF53
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListCategoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCategoriesResponseBody) GetData() []*Category {
	return s.Data
}

func (s *ListCategoriesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCategoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCategoriesResponseBody) SetData(v []*Category) *ListCategoriesResponseBody {
	s.Data = v
	return s
}

func (s *ListCategoriesResponseBody) SetMessage(v string) *ListCategoriesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCategoriesResponseBody) SetRequestId(v string) *ListCategoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCategoriesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
