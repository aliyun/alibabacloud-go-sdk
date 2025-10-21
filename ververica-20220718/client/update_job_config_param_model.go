// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobConfigParam interface {
	dara.Model
	String() string
	GoString() string
	SetNewFlinkConf(v map[string]interface{}) *UpdateJobConfigParam
	GetNewFlinkConf() map[string]interface{}
}

type UpdateJobConfigParam struct {
	NewFlinkConf map[string]interface{} `json:"newFlinkConf,omitempty" xml:"newFlinkConf,omitempty"`
}

func (s UpdateJobConfigParam) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobConfigParam) GoString() string {
	return s.String()
}

func (s *UpdateJobConfigParam) GetNewFlinkConf() map[string]interface{} {
	return s.NewFlinkConf
}

func (s *UpdateJobConfigParam) SetNewFlinkConf(v map[string]interface{}) *UpdateJobConfigParam {
	s.NewFlinkConf = v
	return s
}

func (s *UpdateJobConfigParam) Validate() error {
	return dara.Validate(s)
}
