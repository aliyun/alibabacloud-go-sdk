// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlarmMachineCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *GetAlarmMachineCountRequest
	GetFrom() *string
}

type GetAlarmMachineCountRequest struct {
	// The ID of the request source. Set the value to sas.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
}

func (s GetAlarmMachineCountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmMachineCountRequest) GoString() string {
	return s.String()
}

func (s *GetAlarmMachineCountRequest) GetFrom() *string {
	return s.From
}

func (s *GetAlarmMachineCountRequest) SetFrom(v string) *GetAlarmMachineCountRequest {
	s.From = &v
	return s
}

func (s *GetAlarmMachineCountRequest) Validate() error {
	return dara.Validate(s)
}
