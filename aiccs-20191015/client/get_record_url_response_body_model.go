// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecordUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetRecordUrlResponseBodyData) *GetRecordUrlResponseBody
	GetData() *GetRecordUrlResponseBodyData
	SetMessage(v string) *GetRecordUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRecordUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRecordUrlResponseBody
	GetSuccess() *bool
}

type GetRecordUrlResponseBody struct {
	Data *GetRecordUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRecordUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRecordUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecordUrlResponseBody) GetData() *GetRecordUrlResponseBodyData {
	return s.Data
}

func (s *GetRecordUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRecordUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRecordUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRecordUrlResponseBody) SetData(v *GetRecordUrlResponseBodyData) *GetRecordUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetRecordUrlResponseBody) SetMessage(v string) *GetRecordUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetRecordUrlResponseBody) SetRequestId(v string) *GetRecordUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRecordUrlResponseBody) SetSuccess(v bool) *GetRecordUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetRecordUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRecordUrlResponseBodyData struct {
	// example:
	//
	// 1001067****
	Acid *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	// example:
	//
	// http://aliccrec-shvpc.oss-cn-shanghai.aliyuncs.com/accrec_tmp/10010679716-12-01-56.wav?***
	OssLink *string `json:"OssLink,omitempty" xml:"OssLink,omitempty"`
}

func (s GetRecordUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRecordUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRecordUrlResponseBodyData) GetAcid() *string {
	return s.Acid
}

func (s *GetRecordUrlResponseBodyData) GetOssLink() *string {
	return s.OssLink
}

func (s *GetRecordUrlResponseBodyData) SetAcid(v string) *GetRecordUrlResponseBodyData {
	s.Acid = &v
	return s
}

func (s *GetRecordUrlResponseBodyData) SetOssLink(v string) *GetRecordUrlResponseBodyData {
	s.OssLink = &v
	return s
}

func (s *GetRecordUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
