// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushingTarget interface {
	dara.Model
	String() string
	GoString() string
	SetArn(v string) *PushingTarget
	GetArn() *string
	SetCreateTime(v string) *PushingTarget
	GetCreateTime() *string
	SetDescription(v string) *PushingTarget
	GetDescription() *string
	SetHttpRequestTarget(v *PushingTargetHttpRequestTarget) *PushingTarget
	GetHttpRequestTarget() *PushingTargetHttpRequestTarget
	SetName(v string) *PushingTarget
	GetName() *string
	SetRange(v string) *PushingTarget
	GetRange() *string
	SetTemplateUuid(v string) *PushingTarget
	GetTemplateUuid() *string
	SetType(v string) *PushingTarget
	GetType() *string
	SetUpdateTime(v string) *PushingTarget
	GetUpdateTime() *string
	SetUserId(v string) *PushingTarget
	GetUserId() *string
	SetUuid(v string) *PushingTarget
	GetUuid() *string
}

type PushingTarget struct {
	Arn               *string                         `json:"Arn,omitempty" xml:"Arn,omitempty"`
	CreateTime        *string                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description       *string                         `json:"Description,omitempty" xml:"Description,omitempty"`
	HttpRequestTarget *PushingTargetHttpRequestTarget `json:"HttpRequestTarget,omitempty" xml:"HttpRequestTarget,omitempty" type:"Struct"`
	// This parameter is required.
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Range        *string `json:"Range,omitempty" xml:"Range,omitempty"`
	TemplateUuid *string `json:"TemplateUuid,omitempty" xml:"TemplateUuid,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime   *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Uuid         *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s PushingTarget) String() string {
	return dara.Prettify(s)
}

func (s PushingTarget) GoString() string {
	return s.String()
}

func (s *PushingTarget) GetArn() *string {
	return s.Arn
}

func (s *PushingTarget) GetCreateTime() *string {
	return s.CreateTime
}

func (s *PushingTarget) GetDescription() *string {
	return s.Description
}

func (s *PushingTarget) GetHttpRequestTarget() *PushingTargetHttpRequestTarget {
	return s.HttpRequestTarget
}

func (s *PushingTarget) GetName() *string {
	return s.Name
}

func (s *PushingTarget) GetRange() *string {
	return s.Range
}

func (s *PushingTarget) GetTemplateUuid() *string {
	return s.TemplateUuid
}

func (s *PushingTarget) GetType() *string {
	return s.Type
}

func (s *PushingTarget) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *PushingTarget) GetUserId() *string {
	return s.UserId
}

func (s *PushingTarget) GetUuid() *string {
	return s.Uuid
}

func (s *PushingTarget) SetArn(v string) *PushingTarget {
	s.Arn = &v
	return s
}

func (s *PushingTarget) SetCreateTime(v string) *PushingTarget {
	s.CreateTime = &v
	return s
}

func (s *PushingTarget) SetDescription(v string) *PushingTarget {
	s.Description = &v
	return s
}

func (s *PushingTarget) SetHttpRequestTarget(v *PushingTargetHttpRequestTarget) *PushingTarget {
	s.HttpRequestTarget = v
	return s
}

func (s *PushingTarget) SetName(v string) *PushingTarget {
	s.Name = &v
	return s
}

func (s *PushingTarget) SetRange(v string) *PushingTarget {
	s.Range = &v
	return s
}

func (s *PushingTarget) SetTemplateUuid(v string) *PushingTarget {
	s.TemplateUuid = &v
	return s
}

func (s *PushingTarget) SetType(v string) *PushingTarget {
	s.Type = &v
	return s
}

func (s *PushingTarget) SetUpdateTime(v string) *PushingTarget {
	s.UpdateTime = &v
	return s
}

func (s *PushingTarget) SetUserId(v string) *PushingTarget {
	s.UserId = &v
	return s
}

func (s *PushingTarget) SetUuid(v string) *PushingTarget {
	s.Uuid = &v
	return s
}

func (s *PushingTarget) Validate() error {
	return dara.Validate(s)
}

type PushingTargetHttpRequestTarget struct {
	ContentType         *string                                  `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	EncryptSignatureKey *string                                  `json:"EncryptSignatureKey,omitempty" xml:"EncryptSignatureKey,omitempty"`
	EncryptString       *string                                  `json:"EncryptString,omitempty" xml:"EncryptString,omitempty"`
	EncryptTimestampKey *string                                  `json:"EncryptTimestampKey,omitempty" xml:"EncryptTimestampKey,omitempty"`
	Headers             []*PushingTargetHttpRequestTargetHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	Method              *string                                  `json:"Method,omitempty" xml:"Method,omitempty"`
	Url                 *string                                  `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s PushingTargetHttpRequestTarget) String() string {
	return dara.Prettify(s)
}

func (s PushingTargetHttpRequestTarget) GoString() string {
	return s.String()
}

func (s *PushingTargetHttpRequestTarget) GetContentType() *string {
	return s.ContentType
}

func (s *PushingTargetHttpRequestTarget) GetEncryptSignatureKey() *string {
	return s.EncryptSignatureKey
}

func (s *PushingTargetHttpRequestTarget) GetEncryptString() *string {
	return s.EncryptString
}

func (s *PushingTargetHttpRequestTarget) GetEncryptTimestampKey() *string {
	return s.EncryptTimestampKey
}

func (s *PushingTargetHttpRequestTarget) GetHeaders() []*PushingTargetHttpRequestTargetHeaders {
	return s.Headers
}

func (s *PushingTargetHttpRequestTarget) GetMethod() *string {
	return s.Method
}

func (s *PushingTargetHttpRequestTarget) GetUrl() *string {
	return s.Url
}

func (s *PushingTargetHttpRequestTarget) SetContentType(v string) *PushingTargetHttpRequestTarget {
	s.ContentType = &v
	return s
}

func (s *PushingTargetHttpRequestTarget) SetEncryptSignatureKey(v string) *PushingTargetHttpRequestTarget {
	s.EncryptSignatureKey = &v
	return s
}

func (s *PushingTargetHttpRequestTarget) SetEncryptString(v string) *PushingTargetHttpRequestTarget {
	s.EncryptString = &v
	return s
}

func (s *PushingTargetHttpRequestTarget) SetEncryptTimestampKey(v string) *PushingTargetHttpRequestTarget {
	s.EncryptTimestampKey = &v
	return s
}

func (s *PushingTargetHttpRequestTarget) SetHeaders(v []*PushingTargetHttpRequestTargetHeaders) *PushingTargetHttpRequestTarget {
	s.Headers = v
	return s
}

func (s *PushingTargetHttpRequestTarget) SetMethod(v string) *PushingTargetHttpRequestTarget {
	s.Method = &v
	return s
}

func (s *PushingTargetHttpRequestTarget) SetUrl(v string) *PushingTargetHttpRequestTarget {
	s.Url = &v
	return s
}

func (s *PushingTargetHttpRequestTarget) Validate() error {
	return dara.Validate(s)
}

type PushingTargetHttpRequestTargetHeaders struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PushingTargetHttpRequestTargetHeaders) String() string {
	return dara.Prettify(s)
}

func (s PushingTargetHttpRequestTargetHeaders) GoString() string {
	return s.String()
}

func (s *PushingTargetHttpRequestTargetHeaders) GetKey() *string {
	return s.Key
}

func (s *PushingTargetHttpRequestTargetHeaders) GetValue() *string {
	return s.Value
}

func (s *PushingTargetHttpRequestTargetHeaders) SetKey(v string) *PushingTargetHttpRequestTargetHeaders {
	s.Key = &v
	return s
}

func (s *PushingTargetHttpRequestTargetHeaders) SetValue(v string) *PushingTargetHttpRequestTargetHeaders {
	s.Value = &v
	return s
}

func (s *PushingTargetHttpRequestTargetHeaders) Validate() error {
	return dara.Validate(s)
}
