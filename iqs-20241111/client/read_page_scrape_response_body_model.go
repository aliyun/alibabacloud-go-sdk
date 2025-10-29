// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadPageScrapeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ReadPageItem) *ReadPageScrapeResponseBody
	GetData() *ReadPageItem
	SetErrorCode(v string) *ReadPageScrapeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ReadPageScrapeResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ReadPageScrapeResponseBody
	GetRequestId() *string
}

type ReadPageScrapeResponseBody struct {
	Data *ReadPageItem `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// null
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// null
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 7cd43c86-731a-4d4c-8385-d070cfc509a4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ReadPageScrapeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadPageScrapeResponseBody) GoString() string {
	return s.String()
}

func (s *ReadPageScrapeResponseBody) GetData() *ReadPageItem {
	return s.Data
}

func (s *ReadPageScrapeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ReadPageScrapeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ReadPageScrapeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadPageScrapeResponseBody) SetData(v *ReadPageItem) *ReadPageScrapeResponseBody {
	s.Data = v
	return s
}

func (s *ReadPageScrapeResponseBody) SetErrorCode(v string) *ReadPageScrapeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ReadPageScrapeResponseBody) SetErrorMessage(v string) *ReadPageScrapeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ReadPageScrapeResponseBody) SetRequestId(v string) *ReadPageScrapeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadPageScrapeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
