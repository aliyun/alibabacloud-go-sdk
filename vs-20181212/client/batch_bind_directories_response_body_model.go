// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchBindDirectoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchBindDirectoriesResponseBody
	GetRequestId() *string
	SetResults(v []*BatchBindDirectoriesResponseBodyResults) *BatchBindDirectoriesResponseBody
	GetResults() []*BatchBindDirectoriesResponseBodyResults
}

type BatchBindDirectoriesResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*BatchBindDirectoriesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BatchBindDirectoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchBindDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchBindDirectoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchBindDirectoriesResponseBody) GetResults() []*BatchBindDirectoriesResponseBodyResults {
	return s.Results
}

func (s *BatchBindDirectoriesResponseBody) SetRequestId(v string) *BatchBindDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchBindDirectoriesResponseBody) SetResults(v []*BatchBindDirectoriesResponseBodyResults) *BatchBindDirectoriesResponseBody {
	s.Results = v
	return s
}

func (s *BatchBindDirectoriesResponseBody) Validate() error {
	return dara.Validate(s)
}

type BatchBindDirectoriesResponseBodyResults struct {
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

func (s BatchBindDirectoriesResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s BatchBindDirectoriesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *BatchBindDirectoriesResponseBodyResults) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BatchBindDirectoriesResponseBodyResults) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *BatchBindDirectoriesResponseBodyResults) GetError() *string {
	return s.Error
}

func (s *BatchBindDirectoriesResponseBodyResults) SetDeviceId(v string) *BatchBindDirectoriesResponseBodyResults {
	s.DeviceId = &v
	return s
}

func (s *BatchBindDirectoriesResponseBodyResults) SetDirectoryId(v string) *BatchBindDirectoriesResponseBodyResults {
	s.DirectoryId = &v
	return s
}

func (s *BatchBindDirectoriesResponseBodyResults) SetError(v string) *BatchBindDirectoriesResponseBodyResults {
	s.Error = &v
	return s
}

func (s *BatchBindDirectoriesResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
