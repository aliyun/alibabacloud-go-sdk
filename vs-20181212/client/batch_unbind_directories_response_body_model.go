// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUnbindDirectoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchUnbindDirectoriesResponseBody
	GetRequestId() *string
	SetResults(v []*BatchUnbindDirectoriesResponseBodyResults) *BatchUnbindDirectoriesResponseBody
	GetResults() []*BatchUnbindDirectoriesResponseBodyResults
}

type BatchUnbindDirectoriesResponseBody struct {
	// example:
	//
	// 64DB7F34-11A8-45DC-A421-40ACF446282C
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchUnbindDirectoriesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchUnbindDirectoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchUnbindDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUnbindDirectoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchUnbindDirectoriesResponseBody) GetResults() []*BatchUnbindDirectoriesResponseBodyResults {
	return s.Results
}

func (s *BatchUnbindDirectoriesResponseBody) SetRequestId(v string) *BatchUnbindDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUnbindDirectoriesResponseBody) SetResults(v []*BatchUnbindDirectoriesResponseBodyResults) *BatchUnbindDirectoriesResponseBody {
	s.Results = v
	return s
}

func (s *BatchUnbindDirectoriesResponseBody) Validate() error {
	return dara.Validate(s)
}

type BatchUnbindDirectoriesResponseBodyResults struct {
	// example:
	//
	// 348*****380-cn-qingdao
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// 399*****488-cn-qingdao
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// example:
	//
	// An error occurred while processing your request.
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
}

func (s BatchUnbindDirectoriesResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s BatchUnbindDirectoriesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchUnbindDirectoriesResponseBodyResults) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BatchUnbindDirectoriesResponseBodyResults) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *BatchUnbindDirectoriesResponseBodyResults) GetError() *string {
	return s.Error
}

func (s *BatchUnbindDirectoriesResponseBodyResults) SetDeviceId(v string) *BatchUnbindDirectoriesResponseBodyResults {
	s.DeviceId = &v
	return s
}

func (s *BatchUnbindDirectoriesResponseBodyResults) SetDirectoryId(v string) *BatchUnbindDirectoriesResponseBodyResults {
	s.DirectoryId = &v
	return s
}

func (s *BatchUnbindDirectoriesResponseBodyResults) SetError(v string) *BatchUnbindDirectoriesResponseBodyResults {
	s.Error = &v
	return s
}

func (s *BatchUnbindDirectoriesResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
