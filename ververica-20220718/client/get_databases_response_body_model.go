// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatabasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Database) *GetDatabasesResponseBody
	GetData() []*Database
	SetErrorCode(v string) *GetDatabasesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDatabasesResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetDatabasesResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetDatabasesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDatabasesResponseBody
	GetSuccess() *bool
}

type GetDatabasesResponseBody struct {
	Data []*Database `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDatabasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatabasesResponseBody) GetData() []*Database {
	return s.Data
}

func (s *GetDatabasesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDatabasesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDatabasesResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetDatabasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDatabasesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDatabasesResponseBody) SetData(v []*Database) *GetDatabasesResponseBody {
	s.Data = v
	return s
}

func (s *GetDatabasesResponseBody) SetErrorCode(v string) *GetDatabasesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDatabasesResponseBody) SetErrorMessage(v string) *GetDatabasesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDatabasesResponseBody) SetHttpCode(v int32) *GetDatabasesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetDatabasesResponseBody) SetRequestId(v string) *GetDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatabasesResponseBody) SetSuccess(v bool) *GetDatabasesResponseBody {
	s.Success = &v
	return s
}

func (s *GetDatabasesResponseBody) Validate() error {
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
