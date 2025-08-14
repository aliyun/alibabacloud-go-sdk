// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveAnalysisColumnResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveAnalysisColumnResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *SaveAnalysisColumnResponseBody
	GetResultObject() *bool
}

type SaveAnalysisColumnResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s SaveAnalysisColumnResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveAnalysisColumnResponseBody) GoString() string {
	return s.String()
}

func (s *SaveAnalysisColumnResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveAnalysisColumnResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *SaveAnalysisColumnResponseBody) SetRequestId(v string) *SaveAnalysisColumnResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveAnalysisColumnResponseBody) SetResultObject(v bool) *SaveAnalysisColumnResponseBody {
	s.ResultObject = &v
	return s
}

func (s *SaveAnalysisColumnResponseBody) Validate() error {
	return dara.Validate(s)
}
