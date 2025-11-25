// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceSyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateInstanceSyncTaskRequest
	GetLang() *string
	SetSourceIp(v string) *CreateInstanceSyncTaskRequest
	GetSourceIp() *string
}

type CreateInstanceSyncTaskRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 60.182.79.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s CreateInstanceSyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceSyncTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceSyncTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateInstanceSyncTaskRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *CreateInstanceSyncTaskRequest) SetLang(v string) *CreateInstanceSyncTaskRequest {
	s.Lang = &v
	return s
}

func (s *CreateInstanceSyncTaskRequest) SetSourceIp(v string) *CreateInstanceSyncTaskRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateInstanceSyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
