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
	// The application ID. You can create and query the ID in the console. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The configuration template ID. For more information about how to obtain the template ID, see [AddRecordTemplate](https://help.aliyun.com/document_detail/193625.html).
	//
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
