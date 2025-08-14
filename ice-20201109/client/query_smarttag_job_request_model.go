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
	// The ID of the smart tagging job that you want to query. You can obtain the job ID from the response parameters of the SubmitSmarttagJob operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 88c6ca184c0e47098a5b665e2****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The extra parameters that you want to query in the request. The value is a JSON string. Example: {"labelResultType":"auto"}. The value of labelResultType is of the STRING type. Valid values:
	//
	// 	- auto: machine tagging
	//
	// 	- hmi: tagging by human and machine
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
