// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePluginConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPluginConfigID(v int64) *CreatePluginConfigResponseBody
	GetPluginConfigID() *int64
	SetRequestId(v string) *CreatePluginConfigResponseBody
	GetRequestId() *string
}

type CreatePluginConfigResponseBody struct {
	// The plug-in configuration ID.
	//
	// example:
	//
	// 10
	PluginConfigID *int64 `json:"PluginConfigID,omitempty" xml:"PluginConfigID,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 03A3E2F4-6804-5663-9D5D-2EC47A1*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePluginConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePluginConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePluginConfigResponseBody) GetPluginConfigID() *int64 {
	return s.PluginConfigID
}

func (s *CreatePluginConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePluginConfigResponseBody) SetPluginConfigID(v int64) *CreatePluginConfigResponseBody {
	s.PluginConfigID = &v
	return s
}

func (s *CreatePluginConfigResponseBody) SetRequestId(v string) *CreatePluginConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePluginConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
