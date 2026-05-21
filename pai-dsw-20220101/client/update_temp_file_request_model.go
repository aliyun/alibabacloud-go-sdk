// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTempFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGmtExpiredTime(v string) *UpdateTempFileRequest
	GetGmtExpiredTime() *string
	SetStatus(v string) *UpdateTempFileRequest
	GetStatus() *string
}

type UpdateTempFileRequest struct {
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtExpiredTime *string `json:"GmtExpiredTime,omitempty" xml:"GmtExpiredTime,omitempty"`
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateTempFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTempFileRequest) GoString() string {
	return s.String()
}

func (s *UpdateTempFileRequest) GetGmtExpiredTime() *string {
	return s.GmtExpiredTime
}

func (s *UpdateTempFileRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateTempFileRequest) SetGmtExpiredTime(v string) *UpdateTempFileRequest {
	s.GmtExpiredTime = &v
	return s
}

func (s *UpdateTempFileRequest) SetStatus(v string) *UpdateTempFileRequest {
	s.Status = &v
	return s
}

func (s *UpdateTempFileRequest) Validate() error {
	return dara.Validate(s)
}
