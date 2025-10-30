// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWmEmbedTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateWmEmbedTaskResponseBodyData) *CreateWmEmbedTaskResponseBody
	GetData() *CreateWmEmbedTaskResponseBodyData
	SetRequestId(v string) *CreateWmEmbedTaskResponseBody
	GetRequestId() *string
}

type CreateWmEmbedTaskResponseBody struct {
	Data *CreateWmEmbedTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWmEmbedTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskResponseBody) GetData() *CreateWmEmbedTaskResponseBodyData {
	return s.Data
}

func (s *CreateWmEmbedTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWmEmbedTaskResponseBody) SetData(v *CreateWmEmbedTaskResponseBodyData) *CreateWmEmbedTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateWmEmbedTaskResponseBody) SetRequestId(v string) *CreateWmEmbedTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWmEmbedTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWmEmbedTaskResponseBodyData struct {
	// example:
	//
	// job:5GfrJYsoaffmCE7Z5bZtjUefzxfd****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateWmEmbedTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateWmEmbedTaskResponseBodyData) SetTaskId(v string) *CreateWmEmbedTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateWmEmbedTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
