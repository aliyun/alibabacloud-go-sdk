// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAirflowLoginTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAirflowId(v string) *CreateAirflowLoginTokenRequest
	GetAirflowId() *string
}

type CreateAirflowLoginTokenRequest struct {
	// The ID of the Airflow instance. You can view the instance ID on the [Airflow Instances](https://help.aliyun.com/document_detail/2881043.html) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// af-b3a7f110a6vmvn7xxxxxx
	AirflowId *string `json:"AirflowId,omitempty" xml:"AirflowId,omitempty"`
}

func (s CreateAirflowLoginTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAirflowLoginTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateAirflowLoginTokenRequest) GetAirflowId() *string {
	return s.AirflowId
}

func (s *CreateAirflowLoginTokenRequest) SetAirflowId(v string) *CreateAirflowLoginTokenRequest {
	s.AirflowId = &v
	return s
}

func (s *CreateAirflowLoginTokenRequest) Validate() error {
	return dara.Validate(s)
}
