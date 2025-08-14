// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelSampleDownloadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModelSampleDownloadResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *ModelSampleDownloadResponseBody
	GetResultObject() *bool
}

type ModelSampleDownloadResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Request result.
	//
	// example:
	//
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
}

func (s ModelSampleDownloadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelSampleDownloadResponseBody) GoString() string {
	return s.String()
}

func (s *ModelSampleDownloadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelSampleDownloadResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *ModelSampleDownloadResponseBody) SetRequestId(v string) *ModelSampleDownloadResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelSampleDownloadResponseBody) SetResultObject(v bool) *ModelSampleDownloadResponseBody {
	s.ResultObject = &v
	return s
}

func (s *ModelSampleDownloadResponseBody) Validate() error {
	return dara.Validate(s)
}
