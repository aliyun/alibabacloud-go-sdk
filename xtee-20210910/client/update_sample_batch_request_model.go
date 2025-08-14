// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSampleBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UpdateSampleBatchRequest
	GetLang() *string
	SetIds(v string) *UpdateSampleBatchRequest
	GetIds() *string
	SetRegId(v string) *UpdateSampleBatchRequest
	GetRegId() *string
	SetTags(v string) *UpdateSampleBatchRequest
	GetTags() *string
	SetVersions(v string) *UpdateSampleBatchRequest
	GetVersions() *string
}

type UpdateSampleBatchRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Batch operation IDs.
	//
	// example:
	//
	// 324,343
	Ids *string `json:"ids,omitempty" xml:"ids,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Tags (comma-separated).
	//
	// example:
	//
	// rm0102,rm0103
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// Version number (latest).
	//
	// example:
	//
	// 1,1
	Versions *string `json:"versions,omitempty" xml:"versions,omitempty"`
}

func (s UpdateSampleBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSampleBatchRequest) GoString() string {
	return s.String()
}

func (s *UpdateSampleBatchRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateSampleBatchRequest) GetIds() *string {
	return s.Ids
}

func (s *UpdateSampleBatchRequest) GetRegId() *string {
	return s.RegId
}

func (s *UpdateSampleBatchRequest) GetTags() *string {
	return s.Tags
}

func (s *UpdateSampleBatchRequest) GetVersions() *string {
	return s.Versions
}

func (s *UpdateSampleBatchRequest) SetLang(v string) *UpdateSampleBatchRequest {
	s.Lang = &v
	return s
}

func (s *UpdateSampleBatchRequest) SetIds(v string) *UpdateSampleBatchRequest {
	s.Ids = &v
	return s
}

func (s *UpdateSampleBatchRequest) SetRegId(v string) *UpdateSampleBatchRequest {
	s.RegId = &v
	return s
}

func (s *UpdateSampleBatchRequest) SetTags(v string) *UpdateSampleBatchRequest {
	s.Tags = &v
	return s
}

func (s *UpdateSampleBatchRequest) SetVersions(v string) *UpdateSampleBatchRequest {
	s.Versions = &v
	return s
}

func (s *UpdateSampleBatchRequest) Validate() error {
	return dara.Validate(s)
}
