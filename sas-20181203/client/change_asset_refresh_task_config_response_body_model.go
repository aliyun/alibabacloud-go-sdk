// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeAssetRefreshTaskConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ChangeAssetRefreshTaskConfigResponseBody
	GetData() *bool
	SetMessage(v string) *ChangeAssetRefreshTaskConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChangeAssetRefreshTaskConfigResponseBody
	GetRequestId() *string
}

type ChangeAssetRefreshTaskConfigResponseBody struct {
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BE120DAB-F4E7-4C53-ADC3-A97578ABF384
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeAssetRefreshTaskConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeAssetRefreshTaskConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeAssetRefreshTaskConfigResponseBody) GetData() *bool {
	return s.Data
}

func (s *ChangeAssetRefreshTaskConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChangeAssetRefreshTaskConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeAssetRefreshTaskConfigResponseBody) SetData(v bool) *ChangeAssetRefreshTaskConfigResponseBody {
	s.Data = &v
	return s
}

func (s *ChangeAssetRefreshTaskConfigResponseBody) SetMessage(v string) *ChangeAssetRefreshTaskConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeAssetRefreshTaskConfigResponseBody) SetRequestId(v string) *ChangeAssetRefreshTaskConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeAssetRefreshTaskConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
