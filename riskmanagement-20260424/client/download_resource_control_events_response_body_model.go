// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadResourceControlEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DownloadResourceControlEventsResponseBody
	GetCode() *string
	SetData(v string) *DownloadResourceControlEventsResponseBody
	GetData() *string
	SetMessage(v string) *DownloadResourceControlEventsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DownloadResourceControlEventsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DownloadResourceControlEventsResponseBody
	GetSuccess() *bool
}

type DownloadResourceControlEventsResponseBody struct {
	// The status code.
	//
	// > 200 indicates success. Other values (such as 500 or 400) indicate error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The metadata response information.
	//
	// example:
	//
	// 5
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The prompt message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 855FCC89-0B13-5FC0-AAD2-120878081C1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DownloadResourceControlEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadResourceControlEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadResourceControlEventsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DownloadResourceControlEventsResponseBody) GetData() *string {
	return s.Data
}

func (s *DownloadResourceControlEventsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DownloadResourceControlEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadResourceControlEventsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DownloadResourceControlEventsResponseBody) SetCode(v string) *DownloadResourceControlEventsResponseBody {
	s.Code = &v
	return s
}

func (s *DownloadResourceControlEventsResponseBody) SetData(v string) *DownloadResourceControlEventsResponseBody {
	s.Data = &v
	return s
}

func (s *DownloadResourceControlEventsResponseBody) SetMessage(v string) *DownloadResourceControlEventsResponseBody {
	s.Message = &v
	return s
}

func (s *DownloadResourceControlEventsResponseBody) SetRequestId(v string) *DownloadResourceControlEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadResourceControlEventsResponseBody) SetSuccess(v bool) *DownloadResourceControlEventsResponseBody {
	s.Success = &v
	return s
}

func (s *DownloadResourceControlEventsResponseBody) Validate() error {
	return dara.Validate(s)
}
