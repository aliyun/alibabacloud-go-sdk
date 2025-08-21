// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchDeleteDevicesResponseBody
	GetRequestId() *string
	SetResults(v []*BatchDeleteDevicesResponseBodyResults) *BatchDeleteDevicesResponseBody
	GetResults() []*BatchDeleteDevicesResponseBodyResults
}

type BatchDeleteDevicesResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchDeleteDevicesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchDeleteDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteDevicesResponseBody) GetResults() []*BatchDeleteDevicesResponseBodyResults {
	return s.Results
}

func (s *BatchDeleteDevicesResponseBody) SetRequestId(v string) *BatchDeleteDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteDevicesResponseBody) SetResults(v []*BatchDeleteDevicesResponseBodyResults) *BatchDeleteDevicesResponseBody {
	s.Results = v
	return s
}

func (s *BatchDeleteDevicesResponseBody) Validate() error {
	return dara.Validate(s)
}

type BatchDeleteDevicesResponseBodyResults struct {
	// example:
	//
	// Device not found
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 323884****9092996
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s BatchDeleteDevicesResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDevicesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchDeleteDevicesResponseBodyResults) GetError() *string {
	return s.Error
}

func (s *BatchDeleteDevicesResponseBodyResults) GetId() *string {
	return s.Id
}

func (s *BatchDeleteDevicesResponseBodyResults) SetError(v string) *BatchDeleteDevicesResponseBodyResults {
	s.Error = &v
	return s
}

func (s *BatchDeleteDevicesResponseBodyResults) SetId(v string) *BatchDeleteDevicesResponseBodyResults {
	s.Id = &v
	return s
}

func (s *BatchDeleteDevicesResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
