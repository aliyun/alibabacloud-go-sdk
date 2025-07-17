// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAddonReleaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteAddonReleaseResponseBody
	GetCode() *int32
	SetData(v string) *DeleteAddonReleaseResponseBody
	GetData() *string
	SetMessage(v string) *DeleteAddonReleaseResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAddonReleaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAddonReleaseResponseBody
	GetSuccess() *bool
}

type DeleteAddonReleaseResponseBody struct {
	// Status code: 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return a message.
	//
	// example:
	//
	// success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F7781D4A-2818-41E7-B7BB-79D809E9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the alert rule was deleted. Valid values:
	//
	// 	- `true`: The alert rule was deleted.
	//
	// 	- `false`: The alert rule failed to be deleted.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAddonReleaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAddonReleaseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAddonReleaseResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteAddonReleaseResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteAddonReleaseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAddonReleaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAddonReleaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAddonReleaseResponseBody) SetCode(v int32) *DeleteAddonReleaseResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAddonReleaseResponseBody) SetData(v string) *DeleteAddonReleaseResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAddonReleaseResponseBody) SetMessage(v string) *DeleteAddonReleaseResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAddonReleaseResponseBody) SetRequestId(v string) *DeleteAddonReleaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAddonReleaseResponseBody) SetSuccess(v bool) *DeleteAddonReleaseResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAddonReleaseResponseBody) Validate() error {
	return dara.Validate(s)
}
