// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTestResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecutor(v string) *UpdateTestResultRequest
	GetExecutor() *string
	SetStatus(v string) *UpdateTestResultRequest
	GetStatus() *string
}

type UpdateTestResultRequest struct {
	// example:
	//
	// 131xxx38624xxxx68
	Executor *string `json:"executor,omitempty" xml:"executor,omitempty"`
	// example:
	//
	// TO DO
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateTestResultRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTestResultRequest) GoString() string {
	return s.String()
}

func (s *UpdateTestResultRequest) GetExecutor() *string {
	return s.Executor
}

func (s *UpdateTestResultRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateTestResultRequest) SetExecutor(v string) *UpdateTestResultRequest {
	s.Executor = &v
	return s
}

func (s *UpdateTestResultRequest) SetStatus(v string) *UpdateTestResultRequest {
	s.Status = &v
	return s
}

func (s *UpdateTestResultRequest) Validate() error {
	return dara.Validate(s)
}
