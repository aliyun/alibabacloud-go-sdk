// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSavepointStatus interface {
	dara.Model
	String() string
	GoString() string
	SetFailure(v *SavepointFailure) *SavepointStatus
	GetFailure() *SavepointFailure
	SetState(v string) *SavepointStatus
	GetState() *string
}

type SavepointStatus struct {
	// The details of the failure to create a savepoint for the deployment.
	Failure *SavepointFailure `json:"failure,omitempty" xml:"failure,omitempty"`
	// The status of the savepoint that is created for the deployment. Valid values:
	//
	// 	- STARTED
	//
	// 	- COMPLETED
	//
	// 	- FAILED
	//
	// example:
	//
	// COMPLETED
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s SavepointStatus) String() string {
	return dara.Prettify(s)
}

func (s SavepointStatus) GoString() string {
	return s.String()
}

func (s *SavepointStatus) GetFailure() *SavepointFailure {
	return s.Failure
}

func (s *SavepointStatus) GetState() *string {
	return s.State
}

func (s *SavepointStatus) SetFailure(v *SavepointFailure) *SavepointStatus {
	s.Failure = v
	return s
}

func (s *SavepointStatus) SetState(v string) *SavepointStatus {
	s.State = &v
	return s
}

func (s *SavepointStatus) Validate() error {
	if s.Failure != nil {
		if err := s.Failure.Validate(); err != nil {
			return err
		}
	}
	return nil
}
