// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecordTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteRecordTemplateRequest
	GetAppId() *string
	SetOwnerId(v int64) *DeleteRecordTemplateRequest
	GetOwnerId() *int64
	SetTemplateId(v string) *DeleteRecordTemplateRequest
	GetTemplateId() *string
}

type DeleteRecordTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 1
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 76dasgb****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteRecordTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecordTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecordTemplateRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteRecordTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteRecordTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteRecordTemplateRequest) SetAppId(v string) *DeleteRecordTemplateRequest {
	s.AppId = &v
	return s
}

func (s *DeleteRecordTemplateRequest) SetOwnerId(v int64) *DeleteRecordTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteRecordTemplateRequest) SetTemplateId(v string) *DeleteRecordTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteRecordTemplateRequest) Validate() error {
	return dara.Validate(s)
}
