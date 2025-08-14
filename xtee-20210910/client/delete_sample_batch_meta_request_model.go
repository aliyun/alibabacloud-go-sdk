// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSampleBatchMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteSampleBatchMetaRequest
	GetLang() *string
	SetBatchUuid(v string) *DeleteSampleBatchMetaRequest
	GetBatchUuid() *string
	SetRegId(v string) *DeleteSampleBatchMetaRequest
	GetRegId() *string
}

type DeleteSampleBatchMetaRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// UUID.
	//
	// example:
	//
	// jigaklba83ka
	BatchUuid *string `json:"batchUuid,omitempty" xml:"batchUuid,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DeleteSampleBatchMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSampleBatchMetaRequest) GoString() string {
	return s.String()
}

func (s *DeleteSampleBatchMetaRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteSampleBatchMetaRequest) GetBatchUuid() *string {
	return s.BatchUuid
}

func (s *DeleteSampleBatchMetaRequest) GetRegId() *string {
	return s.RegId
}

func (s *DeleteSampleBatchMetaRequest) SetLang(v string) *DeleteSampleBatchMetaRequest {
	s.Lang = &v
	return s
}

func (s *DeleteSampleBatchMetaRequest) SetBatchUuid(v string) *DeleteSampleBatchMetaRequest {
	s.BatchUuid = &v
	return s
}

func (s *DeleteSampleBatchMetaRequest) SetRegId(v string) *DeleteSampleBatchMetaRequest {
	s.RegId = &v
	return s
}

func (s *DeleteSampleBatchMetaRequest) Validate() error {
	return dara.Validate(s)
}
