// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Table) *GetTablesResponseBody
	GetData() []*Table
	SetErrorCode(v string) *GetTablesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetTablesResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetTablesResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetTablesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTablesResponseBody
	GetSuccess() *bool
}

type GetTablesResponseBody struct {
	// The list and details of tables that meet the conditions when success is true. This value is empty when success is false.
	Data []*Table `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// - If success is false, a business error code is returned.
	//
	// - If success is true, an empty value is returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// - If success is false, a business error message is returned.
	//
	// - If success is true, an empty value is returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The HTTP status code, which is always 200. Use the success field to determine whether the request was successful.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ECE641B2-AB0B-4174-9C3B-885881558637
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTablesResponseBody) GoString() string {
	return s.String()
}

func (s *GetTablesResponseBody) GetData() []*Table {
	return s.Data
}

func (s *GetTablesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTablesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTablesResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTablesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTablesResponseBody) SetData(v []*Table) *GetTablesResponseBody {
	s.Data = v
	return s
}

func (s *GetTablesResponseBody) SetErrorCode(v string) *GetTablesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTablesResponseBody) SetErrorMessage(v string) *GetTablesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetTablesResponseBody) SetHttpCode(v int32) *GetTablesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetTablesResponseBody) SetRequestId(v string) *GetTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTablesResponseBody) SetSuccess(v bool) *GetTablesResponseBody {
	s.Success = &v
	return s
}

func (s *GetTablesResponseBody) Validate() error {
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
