// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadSmapleBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DownloadSmapleBatchRequest
	GetLang() *string
	SetBatchUuid(v string) *DownloadSmapleBatchRequest
	GetBatchUuid() *string
	SetRegId(v string) *DownloadSmapleBatchRequest
	GetRegId() *string
}

type DownloadSmapleBatchRequest struct {
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
	// Sample batch UUID
	//
	// example:
	//
	// 203f1ae65c0a41a49dc4f8a47946caa2
	BatchUuid *string `json:"batchUuid,omitempty" xml:"batchUuid,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DownloadSmapleBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadSmapleBatchRequest) GoString() string {
	return s.String()
}

func (s *DownloadSmapleBatchRequest) GetLang() *string {
	return s.Lang
}

func (s *DownloadSmapleBatchRequest) GetBatchUuid() *string {
	return s.BatchUuid
}

func (s *DownloadSmapleBatchRequest) GetRegId() *string {
	return s.RegId
}

func (s *DownloadSmapleBatchRequest) SetLang(v string) *DownloadSmapleBatchRequest {
	s.Lang = &v
	return s
}

func (s *DownloadSmapleBatchRequest) SetBatchUuid(v string) *DownloadSmapleBatchRequest {
	s.BatchUuid = &v
	return s
}

func (s *DownloadSmapleBatchRequest) SetRegId(v string) *DownloadSmapleBatchRequest {
	s.RegId = &v
	return s
}

func (s *DownloadSmapleBatchRequest) Validate() error {
	return dara.Validate(s)
}
