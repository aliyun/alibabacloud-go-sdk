// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCloudRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopCloudRecordResponseBody
	GetCode() *string
	SetRequestId(v string) *StopCloudRecordResponseBody
	GetRequestId() *string
}

type StopCloudRecordResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StopCloudRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopCloudRecordResponseBody) GoString() string {
	return s.String()
}

func (s *StopCloudRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopCloudRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopCloudRecordResponseBody) SetCode(v string) *StopCloudRecordResponseBody {
	s.Code = &v
	return s
}

func (s *StopCloudRecordResponseBody) SetRequestId(v string) *StopCloudRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopCloudRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
