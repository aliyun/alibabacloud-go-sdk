// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadVerifyRecordIntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DownloadVerifyRecordIntlResponseBody
	GetCode() *string
	SetData(v *DownloadVerifyRecordIntlResponseBodyData) *DownloadVerifyRecordIntlResponseBody
	GetData() *DownloadVerifyRecordIntlResponseBodyData
	SetMessage(v string) *DownloadVerifyRecordIntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *DownloadVerifyRecordIntlResponseBody
	GetRequestId() *string
}

type DownloadVerifyRecordIntlResponseBody struct {
	// Return code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data.
	Data *DownloadVerifyRecordIntlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 86C40EC3-5940-5F47-995C-BFE90B70E540
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DownloadVerifyRecordIntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadVerifyRecordIntlResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadVerifyRecordIntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *DownloadVerifyRecordIntlResponseBody) GetData() *DownloadVerifyRecordIntlResponseBodyData {
	return s.Data
}

func (s *DownloadVerifyRecordIntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DownloadVerifyRecordIntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadVerifyRecordIntlResponseBody) SetCode(v string) *DownloadVerifyRecordIntlResponseBody {
	s.Code = &v
	return s
}

func (s *DownloadVerifyRecordIntlResponseBody) SetData(v *DownloadVerifyRecordIntlResponseBodyData) *DownloadVerifyRecordIntlResponseBody {
	s.Data = v
	return s
}

func (s *DownloadVerifyRecordIntlResponseBody) SetMessage(v string) *DownloadVerifyRecordIntlResponseBody {
	s.Message = &v
	return s
}

func (s *DownloadVerifyRecordIntlResponseBody) SetRequestId(v string) *DownloadVerifyRecordIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadVerifyRecordIntlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DownloadVerifyRecordIntlResponseBodyData struct {
	// Task ID, returned in asynchronous mode, used later with QueryDownloadTaskIntl to download the exported file.
	//
	// example:
	//
	// 202511284106866
	DownloadTaskId *string `json:"DownloadTaskId,omitempty" xml:"DownloadTaskId,omitempty"`
	// Exported file download link.
	//
	// example:
	//
	// https://cn-shanghai-aliyun-cloudauth.oss-cn-shanghai.aliyuncs.com/console/xxxxxx/KYC_INVOKE_STATISTICS.xlxs
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DownloadVerifyRecordIntlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DownloadVerifyRecordIntlResponseBodyData) GoString() string {
	return s.String()
}

func (s *DownloadVerifyRecordIntlResponseBodyData) GetDownloadTaskId() *string {
	return s.DownloadTaskId
}

func (s *DownloadVerifyRecordIntlResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *DownloadVerifyRecordIntlResponseBodyData) SetDownloadTaskId(v string) *DownloadVerifyRecordIntlResponseBodyData {
	s.DownloadTaskId = &v
	return s
}

func (s *DownloadVerifyRecordIntlResponseBodyData) SetUrl(v string) *DownloadVerifyRecordIntlResponseBodyData {
	s.Url = &v
	return s
}

func (s *DownloadVerifyRecordIntlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
