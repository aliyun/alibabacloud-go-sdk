// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCloudRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartCloudRecordResponseBody
	GetCode() *string
	SetRequestId(v string) *StartCloudRecordResponseBody
	GetRequestId() *string
}

type StartCloudRecordResponseBody struct {
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

func (s StartCloudRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordResponseBody) GoString() string {
	return s.String()
}

func (s *StartCloudRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartCloudRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartCloudRecordResponseBody) SetCode(v string) *StartCloudRecordResponseBody {
	s.Code = &v
	return s
}

func (s *StartCloudRecordResponseBody) SetRequestId(v string) *StartCloudRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartCloudRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
