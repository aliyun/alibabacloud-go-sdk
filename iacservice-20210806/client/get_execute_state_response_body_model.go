// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExecuteStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMessage(v string) *GetExecuteStateResponseBody
	GetErrorMessage() *string
	SetLogFile(v map[string]interface{}) *GetExecuteStateResponseBody
	GetLogFile() map[string]interface{}
	SetRequestId(v string) *GetExecuteStateResponseBody
	GetRequestId() *string
	SetState(v string) *GetExecuteStateResponseBody
	GetState() *string
	SetStatus(v string) *GetExecuteStateResponseBody
	GetStatus() *string
}

type GetExecuteStateResponseBody struct {
	// example:
	//
	// Your account does not have enough balance to order postpaid product.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// {"tf-plan.run.log":"xxx"}
	LogFile map[string]interface{} `json:"logFile,omitempty" xml:"logFile,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B4672AE3-C313-5B7A-BB24-45345570D398
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// {"version": 4, "terraform_version": "1.5.7", "serial": 3, "lineage": "cb71b0b2-1ec2-6483-d409-8cae23186ec6",  "outputs": {}, "resources": [], "check_results": null}
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// Pending
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetExecuteStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExecuteStateResponseBody) GoString() string {
	return s.String()
}

func (s *GetExecuteStateResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetExecuteStateResponseBody) GetLogFile() map[string]interface{} {
	return s.LogFile
}

func (s *GetExecuteStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExecuteStateResponseBody) GetState() *string {
	return s.State
}

func (s *GetExecuteStateResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetExecuteStateResponseBody) SetErrorMessage(v string) *GetExecuteStateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetExecuteStateResponseBody) SetLogFile(v map[string]interface{}) *GetExecuteStateResponseBody {
	s.LogFile = v
	return s
}

func (s *GetExecuteStateResponseBody) SetRequestId(v string) *GetExecuteStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExecuteStateResponseBody) SetState(v string) *GetExecuteStateResponseBody {
	s.State = &v
	return s
}

func (s *GetExecuteStateResponseBody) SetStatus(v string) *GetExecuteStateResponseBody {
	s.Status = &v
	return s
}

func (s *GetExecuteStateResponseBody) Validate() error {
	return dara.Validate(s)
}
