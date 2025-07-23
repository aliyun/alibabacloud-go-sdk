// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentJobStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *UpdateDeploymentJobStatusResponseBody
	GetData() interface{}
	SetRequestId(v string) *UpdateDeploymentJobStatusResponseBody
	GetRequestId() *string
}

type UpdateDeploymentJobStatusResponseBody struct {
	// The response parameters.
	//
	// example:
	//
	// []
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EA69E364-5CBB-50E8-BF09-E8CAA396A4F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDeploymentJobStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentJobStatusResponseBody) GetData() interface{} {
	return s.Data
}

func (s *UpdateDeploymentJobStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDeploymentJobStatusResponseBody) SetData(v interface{}) *UpdateDeploymentJobStatusResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDeploymentJobStatusResponseBody) SetRequestId(v string) *UpdateDeploymentJobStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDeploymentJobStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
