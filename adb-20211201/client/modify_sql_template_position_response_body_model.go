// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySqlTemplatePositionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ModifySqlTemplatePositionResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifySqlTemplatePositionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifySqlTemplatePositionResponseBody
	GetSuccess() *bool
}

type ModifySqlTemplatePositionResponseBody struct {
	// The returned message. Valid values:
	//
	// 	- If the request is successful, a **SUCCESS*	- message is returned.
	//
	// 	- If the request is abnormal, the detailed error message is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5DC10091-348D-12B1-906D-AB49D658012E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: Succeeded.
	//
	// 	- **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifySqlTemplatePositionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySqlTemplatePositionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySqlTemplatePositionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifySqlTemplatePositionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySqlTemplatePositionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifySqlTemplatePositionResponseBody) SetMessage(v string) *ModifySqlTemplatePositionResponseBody {
	s.Message = &v
	return s
}

func (s *ModifySqlTemplatePositionResponseBody) SetRequestId(v string) *ModifySqlTemplatePositionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySqlTemplatePositionResponseBody) SetSuccess(v bool) *ModifySqlTemplatePositionResponseBody {
	s.Success = &v
	return s
}

func (s *ModifySqlTemplatePositionResponseBody) Validate() error {
	return dara.Validate(s)
}
