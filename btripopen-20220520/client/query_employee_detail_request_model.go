// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEmployeeDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOutEmployeeId(v string) *QueryEmployeeDetailRequest
	GetOutEmployeeId() *string
}

type QueryEmployeeDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// abc12138
	OutEmployeeId *string `json:"out_employee_id,omitempty" xml:"out_employee_id,omitempty"`
}

func (s QueryEmployeeDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryEmployeeDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryEmployeeDetailRequest) GetOutEmployeeId() *string {
	return s.OutEmployeeId
}

func (s *QueryEmployeeDetailRequest) SetOutEmployeeId(v string) *QueryEmployeeDetailRequest {
	s.OutEmployeeId = &v
	return s
}

func (s *QueryEmployeeDetailRequest) Validate() error {
	return dara.Validate(s)
}
