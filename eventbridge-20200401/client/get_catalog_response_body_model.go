// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCatalogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCatalogResponseBody
	GetCode() *string
	SetData(v *Catalog) *GetCatalogResponseBody
	GetData() *Catalog
	SetMessage(v string) *GetCatalogResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCatalogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCatalogResponseBody
	GetSuccess() *bool
}

type GetCatalogResponseBody struct {
	// example:
	//
	// 200
	Code *string  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *Catalog `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s GetCatalogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *GetCatalogResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCatalogResponseBody) GetData() *Catalog {
	return s.Data
}

func (s *GetCatalogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCatalogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCatalogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCatalogResponseBody) SetCode(v string) *GetCatalogResponseBody {
	s.Code = &v
	return s
}

func (s *GetCatalogResponseBody) SetData(v *Catalog) *GetCatalogResponseBody {
	s.Data = v
	return s
}

func (s *GetCatalogResponseBody) SetMessage(v string) *GetCatalogResponseBody {
	s.Message = &v
	return s
}

func (s *GetCatalogResponseBody) SetRequestId(v string) *GetCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCatalogResponseBody) SetSuccess(v bool) *GetCatalogResponseBody {
	s.Success = &v
	return s
}

func (s *GetCatalogResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
