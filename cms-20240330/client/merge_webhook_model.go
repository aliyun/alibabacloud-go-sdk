// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMergeWebhook interface {
	dara.Model
	String() string
	GoString() string
	SetContentType(v string) *MergeWebhook
	GetContentType() *string
	SetExtend(v string) *MergeWebhook
	GetExtend() *string
	SetGmtCreate(v string) *MergeWebhook
	GetGmtCreate() *string
	SetGmtModified(v map[string]interface{}) *MergeWebhook
	GetGmtModified() map[string]interface{}
	SetHeaders(v string) *MergeWebhook
	GetHeaders() *string
	SetIdentifier(v string) *MergeWebhook
	GetIdentifier() *string
	SetLang(v string) *MergeWebhook
	GetLang() *string
	SetMethod(v string) *MergeWebhook
	GetMethod() *string
	SetName(v string) *MergeWebhook
	GetName() *string
	SetSource(v string) *MergeWebhook
	GetSource() *string
	SetType(v string) *MergeWebhook
	GetType() *string
	SetWebhook(v string) *MergeWebhook
	GetWebhook() *string
}

type MergeWebhook struct {
	ContentType *string                `json:"contentType,omitempty" xml:"contentType,omitempty"`
	Extend      *string                `json:"extend,omitempty" xml:"extend,omitempty"`
	GmtCreate   *string                `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified map[string]interface{} `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Headers     *string                `json:"headers,omitempty" xml:"headers,omitempty"`
	Identifier  *string                `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Lang        *string                `json:"lang,omitempty" xml:"lang,omitempty"`
	Method      *string                `json:"method,omitempty" xml:"method,omitempty"`
	Name        *string                `json:"name,omitempty" xml:"name,omitempty"`
	Source      *string                `json:"source,omitempty" xml:"source,omitempty"`
	Type        *string                `json:"type,omitempty" xml:"type,omitempty"`
	Webhook     *string                `json:"webhook,omitempty" xml:"webhook,omitempty"`
}

func (s MergeWebhook) String() string {
	return dara.Prettify(s)
}

func (s MergeWebhook) GoString() string {
	return s.String()
}

func (s *MergeWebhook) GetContentType() *string {
	return s.ContentType
}

func (s *MergeWebhook) GetExtend() *string {
	return s.Extend
}

func (s *MergeWebhook) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *MergeWebhook) GetGmtModified() map[string]interface{} {
	return s.GmtModified
}

func (s *MergeWebhook) GetHeaders() *string {
	return s.Headers
}

func (s *MergeWebhook) GetIdentifier() *string {
	return s.Identifier
}

func (s *MergeWebhook) GetLang() *string {
	return s.Lang
}

func (s *MergeWebhook) GetMethod() *string {
	return s.Method
}

func (s *MergeWebhook) GetName() *string {
	return s.Name
}

func (s *MergeWebhook) GetSource() *string {
	return s.Source
}

func (s *MergeWebhook) GetType() *string {
	return s.Type
}

func (s *MergeWebhook) GetWebhook() *string {
	return s.Webhook
}

func (s *MergeWebhook) SetContentType(v string) *MergeWebhook {
	s.ContentType = &v
	return s
}

func (s *MergeWebhook) SetExtend(v string) *MergeWebhook {
	s.Extend = &v
	return s
}

func (s *MergeWebhook) SetGmtCreate(v string) *MergeWebhook {
	s.GmtCreate = &v
	return s
}

func (s *MergeWebhook) SetGmtModified(v map[string]interface{}) *MergeWebhook {
	s.GmtModified = v
	return s
}

func (s *MergeWebhook) SetHeaders(v string) *MergeWebhook {
	s.Headers = &v
	return s
}

func (s *MergeWebhook) SetIdentifier(v string) *MergeWebhook {
	s.Identifier = &v
	return s
}

func (s *MergeWebhook) SetLang(v string) *MergeWebhook {
	s.Lang = &v
	return s
}

func (s *MergeWebhook) SetMethod(v string) *MergeWebhook {
	s.Method = &v
	return s
}

func (s *MergeWebhook) SetName(v string) *MergeWebhook {
	s.Name = &v
	return s
}

func (s *MergeWebhook) SetSource(v string) *MergeWebhook {
	s.Source = &v
	return s
}

func (s *MergeWebhook) SetType(v string) *MergeWebhook {
	s.Type = &v
	return s
}

func (s *MergeWebhook) SetWebhook(v string) *MergeWebhook {
	s.Webhook = &v
	return s
}

func (s *MergeWebhook) Validate() error {
	return dara.Validate(s)
}
