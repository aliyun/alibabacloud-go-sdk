// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteSampleDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *BatchDeleteSampleDataRequest
	GetLang() *string
	SetRegId(v string) *BatchDeleteSampleDataRequest
	GetRegId() *string
	SetUuids(v string) *BatchDeleteSampleDataRequest
	GetUuids() *string
}

type BatchDeleteSampleDataRequest struct {
	// Language of the error message returned by the interface. Values: zh: Chinese; en: English. Default is en.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// UUID.
	//
	// example:
	//
	// 89cd3e44cd2f4a529fb020f3bab3ee1c
	Uuids *string `json:"uuids,omitempty" xml:"uuids,omitempty"`
}

func (s BatchDeleteSampleDataRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteSampleDataRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteSampleDataRequest) GetLang() *string {
	return s.Lang
}

func (s *BatchDeleteSampleDataRequest) GetRegId() *string {
	return s.RegId
}

func (s *BatchDeleteSampleDataRequest) GetUuids() *string {
	return s.Uuids
}

func (s *BatchDeleteSampleDataRequest) SetLang(v string) *BatchDeleteSampleDataRequest {
	s.Lang = &v
	return s
}

func (s *BatchDeleteSampleDataRequest) SetRegId(v string) *BatchDeleteSampleDataRequest {
	s.RegId = &v
	return s
}

func (s *BatchDeleteSampleDataRequest) SetUuids(v string) *BatchDeleteSampleDataRequest {
	s.Uuids = &v
	return s
}

func (s *BatchDeleteSampleDataRequest) Validate() error {
	return dara.Validate(s)
}
