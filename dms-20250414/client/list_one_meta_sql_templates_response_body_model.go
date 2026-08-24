// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOneMetaSqlTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*OneMetaSqlTemplateView) *ListOneMetaSqlTemplatesResponseBody
	GetData() []*OneMetaSqlTemplateView
	SetErrorCode(v string) *ListOneMetaSqlTemplatesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListOneMetaSqlTemplatesResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *ListOneMetaSqlTemplatesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListOneMetaSqlTemplatesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListOneMetaSqlTemplatesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListOneMetaSqlTemplatesResponseBody
	GetSuccess() *bool
}

type ListOneMetaSqlTemplatesResponseBody struct {
	Data         []*OneMetaSqlTemplateView `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode    *string                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	MaxResults   *int32                    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListOneMetaSqlTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOneMetaSqlTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListOneMetaSqlTemplatesResponseBody) GetData() []*OneMetaSqlTemplateView {
	return s.Data
}

func (s *ListOneMetaSqlTemplatesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListOneMetaSqlTemplatesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListOneMetaSqlTemplatesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOneMetaSqlTemplatesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOneMetaSqlTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOneMetaSqlTemplatesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListOneMetaSqlTemplatesResponseBody) SetData(v []*OneMetaSqlTemplateView) *ListOneMetaSqlTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListOneMetaSqlTemplatesResponseBody) SetErrorCode(v string) *ListOneMetaSqlTemplatesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListOneMetaSqlTemplatesResponseBody) SetErrorMessage(v string) *ListOneMetaSqlTemplatesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListOneMetaSqlTemplatesResponseBody) SetMaxResults(v int32) *ListOneMetaSqlTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListOneMetaSqlTemplatesResponseBody) SetNextToken(v string) *ListOneMetaSqlTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOneMetaSqlTemplatesResponseBody) SetRequestId(v string) *ListOneMetaSqlTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOneMetaSqlTemplatesResponseBody) SetSuccess(v bool) *ListOneMetaSqlTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *ListOneMetaSqlTemplatesResponseBody) Validate() error {
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
