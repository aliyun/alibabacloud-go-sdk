// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLumaCatalogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetLumaCatalogResponseBody
	GetCode() *string
	SetData(v *Catalog) *GetLumaCatalogResponseBody
	GetData() *Catalog
	SetMessage(v string) *GetLumaCatalogResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLumaCatalogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLumaCatalogResponseBody
	GetSuccess() *bool
}

type GetLumaCatalogResponseBody struct {
	// The response code. A value of Success indicates a successful call. A specific error code is returned if the call fails.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the data catalog bound to the Agent.
	Data *Catalog `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned by the operation. The value is Operation success when the call succeeds, or a specific error description when the call fails.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, used for troubleshooting and ticket submission.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. A value of true indicates success.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLumaCatalogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLumaCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *GetLumaCatalogResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetLumaCatalogResponseBody) GetData() *Catalog {
	return s.Data
}

func (s *GetLumaCatalogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLumaCatalogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLumaCatalogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLumaCatalogResponseBody) SetCode(v string) *GetLumaCatalogResponseBody {
	s.Code = &v
	return s
}

func (s *GetLumaCatalogResponseBody) SetData(v *Catalog) *GetLumaCatalogResponseBody {
	s.Data = v
	return s
}

func (s *GetLumaCatalogResponseBody) SetMessage(v string) *GetLumaCatalogResponseBody {
	s.Message = &v
	return s
}

func (s *GetLumaCatalogResponseBody) SetRequestId(v string) *GetLumaCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLumaCatalogResponseBody) SetSuccess(v bool) *GetLumaCatalogResponseBody {
	s.Success = &v
	return s
}

func (s *GetLumaCatalogResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
