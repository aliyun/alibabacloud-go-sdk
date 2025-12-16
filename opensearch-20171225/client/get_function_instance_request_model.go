// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOutput(v string) *GetFunctionInstanceRequest
	GetOutput() *string
}

type GetFunctionInstanceRequest struct {
	// Specifies the richness of returned information. Valid values:
	//
	// 	- simple: displays only the basic information.
	//
	// 	- normal: displays information such as createParameters and cron. This is the default value.
	//
	// 	- detail: returns the details of the training task.
	//
	// example:
	//
	// detail
	Output *string `json:"output,omitempty" xml:"output,omitempty"`
}

func (s GetFunctionInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceRequest) GetOutput() *string {
	return s.Output
}

func (s *GetFunctionInstanceRequest) SetOutput(v string) *GetFunctionInstanceRequest {
	s.Output = &v
	return s
}

func (s *GetFunctionInstanceRequest) Validate() error {
	return dara.Validate(s)
}
