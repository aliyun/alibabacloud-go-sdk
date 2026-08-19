// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCrawlerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *StopCrawlerResponseBody
	GetId() *int64
	SetRequestId(v string) *StopCrawlerResponseBody
	GetRequestId() *string
	SetStopAccepted(v bool) *StopCrawlerResponseBody
	GetStopAccepted() *bool
	SetSuccess(v bool) *StopCrawlerResponseBody
	GetSuccess() *bool
}

type StopCrawlerResponseBody struct {
	// The ID of the metadata crawler.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID. Used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 9252F32F-D855-549E-8898-61CF5A733050
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the stop request was accepted. A value of true indicates that the request was accepted. Call ListCrawlerRuns to confirm the final run status.
	StopAccepted *bool `json:"StopAccepted,omitempty" xml:"StopAccepted,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopCrawlerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopCrawlerResponseBody) GoString() string {
	return s.String()
}

func (s *StopCrawlerResponseBody) GetId() *int64 {
	return s.Id
}

func (s *StopCrawlerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopCrawlerResponseBody) GetStopAccepted() *bool {
	return s.StopAccepted
}

func (s *StopCrawlerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopCrawlerResponseBody) SetId(v int64) *StopCrawlerResponseBody {
	s.Id = &v
	return s
}

func (s *StopCrawlerResponseBody) SetRequestId(v string) *StopCrawlerResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopCrawlerResponseBody) SetStopAccepted(v bool) *StopCrawlerResponseBody {
	s.StopAccepted = &v
	return s
}

func (s *StopCrawlerResponseBody) SetSuccess(v bool) *StopCrawlerResponseBody {
	s.Success = &v
	return s
}

func (s *StopCrawlerResponseBody) Validate() error {
	return dara.Validate(s)
}
