// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVodTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddVodTemplateResponseBody
	GetRequestId() *string
	SetVodTemplateId(v string) *AddVodTemplateResponseBody
	GetVodTemplateId() *string
}

type AddVodTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the snapshot template. You can call the [SubmitSnapshotJob](https://help.aliyun.com/document_detail/72213.html) operation to submit snapshot jobs.
	//
	// example:
	//
	// f5b228fe6930e*****0d6bf55bd87789
	VodTemplateId *string `json:"VodTemplateId,omitempty" xml:"VodTemplateId,omitempty"`
}

func (s AddVodTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddVodTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *AddVodTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddVodTemplateResponseBody) GetVodTemplateId() *string {
	return s.VodTemplateId
}

func (s *AddVodTemplateResponseBody) SetRequestId(v string) *AddVodTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddVodTemplateResponseBody) SetVodTemplateId(v string) *AddVodTemplateResponseBody {
	s.VodTemplateId = &v
	return s
}

func (s *AddVodTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
