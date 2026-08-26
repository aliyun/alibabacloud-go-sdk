// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOneMetaOssieModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*OssieModelView) *ListOneMetaOssieModelsResponseBody
	GetData() []*OssieModelView
	SetErrorCode(v string) *ListOneMetaOssieModelsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListOneMetaOssieModelsResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *ListOneMetaOssieModelsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListOneMetaOssieModelsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListOneMetaOssieModelsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListOneMetaOssieModelsResponseBody
	GetSuccess() *bool
}

type ListOneMetaOssieModelsResponseBody struct {
	// The response struct.
	Data []*OssieModelView `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code returned when the request is abnormal.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the call failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The page size.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query.
	//
	// example:
	//
	// NesLoKLEdIZrKhDT7I2gS****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListOneMetaOssieModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOneMetaOssieModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOneMetaOssieModelsResponseBody) GetData() []*OssieModelView {
	return s.Data
}

func (s *ListOneMetaOssieModelsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListOneMetaOssieModelsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListOneMetaOssieModelsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOneMetaOssieModelsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOneMetaOssieModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOneMetaOssieModelsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListOneMetaOssieModelsResponseBody) SetData(v []*OssieModelView) *ListOneMetaOssieModelsResponseBody {
	s.Data = v
	return s
}

func (s *ListOneMetaOssieModelsResponseBody) SetErrorCode(v string) *ListOneMetaOssieModelsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListOneMetaOssieModelsResponseBody) SetErrorMessage(v string) *ListOneMetaOssieModelsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListOneMetaOssieModelsResponseBody) SetMaxResults(v int32) *ListOneMetaOssieModelsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListOneMetaOssieModelsResponseBody) SetNextToken(v string) *ListOneMetaOssieModelsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOneMetaOssieModelsResponseBody) SetRequestId(v string) *ListOneMetaOssieModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOneMetaOssieModelsResponseBody) SetSuccess(v bool) *ListOneMetaOssieModelsResponseBody {
	s.Success = &v
	return s
}

func (s *ListOneMetaOssieModelsResponseBody) Validate() error {
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
