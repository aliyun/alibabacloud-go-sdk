// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCacheAnalysisTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateCacheAnalysisTaskResponseBody
	GetRequestId() *string
}

type CreateCacheAnalysisTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// BBC1E3D6-7C88-4DF5-9A3D-0DB1E6D9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCacheAnalysisTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCacheAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCacheAnalysisTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCacheAnalysisTaskResponseBody) SetRequestId(v string) *CreateCacheAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCacheAnalysisTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
