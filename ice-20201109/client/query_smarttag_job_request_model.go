// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmarttagJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *QuerySmarttagJobRequest
	GetJobId() *string
	SetParams(v string) *QuerySmarttagJobRequest
	GetParams() *string
}

type QuerySmarttagJobRequest struct {
	// The ID of the smart tagging job. You can obtain this ID from the response to the [SubmitSmarttagJob](https://help.aliyun.com/document_detail/478786.html) call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 88c6ca184c0e47098a5b665e2****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Additional request parameters, formatted as a JSON string. For example: `{"labelResultType":"auto"}`. The `labelResultType` parameter supports the following values:
	//
	// - `auto`: machine-generated tagging results
	//
	// - `hmi`: human-in-the-loop tagging results
	//
	// example:
	//
	// {"labelResultType":"auto"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s QuerySmarttagJobRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySmarttagJobRequest) GoString() string {
	return s.String()
}

func (s *QuerySmarttagJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *QuerySmarttagJobRequest) GetParams() *string {
	return s.Params
}

func (s *QuerySmarttagJobRequest) SetJobId(v string) *QuerySmarttagJobRequest {
	s.JobId = &v
	return s
}

func (s *QuerySmarttagJobRequest) SetParams(v string) *QuerySmarttagJobRequest {
	s.Params = &v
	return s
}

func (s *QuerySmarttagJobRequest) Validate() error {
	return dara.Validate(s)
}
