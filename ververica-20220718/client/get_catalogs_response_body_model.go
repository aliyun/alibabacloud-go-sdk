// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCatalogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Catalog) *GetCatalogsResponseBody
	GetData() []*Catalog
	SetErrorCode(v string) *GetCatalogsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetCatalogsResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetCatalogsResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetCatalogsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCatalogsResponseBody
	GetSuccess() *bool
}

type GetCatalogsResponseBody struct {
	Data []*Catalog `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-ABCD-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetCatalogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogsResponseBody) GoString() string {
	return s.String()
}

func (s *GetCatalogsResponseBody) GetData() []*Catalog {
	return s.Data
}

func (s *GetCatalogsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetCatalogsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetCatalogsResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetCatalogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCatalogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCatalogsResponseBody) SetData(v []*Catalog) *GetCatalogsResponseBody {
	s.Data = v
	return s
}

func (s *GetCatalogsResponseBody) SetErrorCode(v string) *GetCatalogsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetCatalogsResponseBody) SetErrorMessage(v string) *GetCatalogsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetCatalogsResponseBody) SetHttpCode(v int32) *GetCatalogsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetCatalogsResponseBody) SetRequestId(v string) *GetCatalogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCatalogsResponseBody) SetSuccess(v bool) *GetCatalogsResponseBody {
	s.Success = &v
	return s
}

func (s *GetCatalogsResponseBody) Validate() error {
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
