// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDevelop(v string) *CreateServiceShrinkRequest
	GetDevelop() *string
	SetLabelsShrink(v string) *CreateServiceShrinkRequest
	GetLabelsShrink() *string
	SetWorkspaceId(v string) *CreateServiceShrinkRequest
	GetWorkspaceId() *string
	SetBody(v string) *CreateServiceShrinkRequest
	GetBody() *string
}

type CreateServiceShrinkRequest struct {
	// Specifies whether to enter development mode.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Develop *string `json:"Develop,omitempty" xml:"Develop,omitempty"`
	// The custom label.
	LabelsShrink *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 123456
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The request body. For more information about the key request parameters, see **Table 1. Request body parameters*	- and **Table 2. Metadata parameters**. For more information about all related parameters, see [Parameters of model services](https://help.aliyun.com/document_detail/450525.html).
	//
	// example:
	//
	// Service deployment by using an image:
	//
	// {
	//
	//   "name": "foo",
	//
	//   "metadata": {
	//
	//     "instance": 2,
	//
	//     "memory": 7000,
	//
	//     "cpu": 4
	//
	//     },
	//
	//   "containers": [
	//
	//     {
	//
	//       "image": "****",
	//
	//       "script": "***	- --listen=0.0.0.0 --server_port=8000 --headless",
	//
	//       "port": 8000
	//
	//     }
	//
	//   ],
	//
	//   "storage": [
	//
	//     {
	//
	//       "oss": {
	//
	//         "path": "oss://examplebuket/data111/",
	//
	//         "readOnly": false
	//
	//       },
	//
	//       "properties": {
	//
	//         "resource_type": "model"
	//
	//       },
	//
	//       "mount_path": "/data"
	//
	//     }
	//
	//   ]
	//
	// }
	//
	// AI-Web application deployment by using an image:
	//
	// {
	//
	//   "name": "foo",
	//
	//   "metadata": {
	//
	//     "instance": 1,
	//
	//     "memory": 7000,
	//
	//     "cpu": 4,
	//
	//     "enable_webservice": true
	//
	//   },
	//
	//   "containers": [
	//
	//     {
	//
	//       "image": "****",
	//
	//       "script": "***	- --listen=0.0.0.0 --server_port=8000 --headless",
	//
	//       "port": 8000
	//
	//     }
	//
	//   ],
	//
	//   "storage": [
	//
	//     {
	//
	//       "oss": {
	//
	//         "path": "oss://examplebucket/data111/",
	//
	//         "readOnly": false
	//
	//       },
	//
	//       "properties": {
	//
	//       "resource_type": "model"
	//
	//       },
	//
	//       "mount_path": "/data"
	//
	//     }
	//
	//   ]
	//
	// }
	//
	// Service deployment by using models and processors:
	//
	// {
	//
	//   "metadata": {
	//
	//     "instance": 1,
	//
	//     "memory": 7000,
	//
	//     "cpu": 4
	//
	//   },
	//
	//   "name": "foo",
	//
	//   "model_config": {},
	//
	//   "processor_type": "python",
	//
	//   "processor_path": "oss://****",
	//
	//   "processor_entry": "a.py",
	//
	//   "model_path": "oss://****"
	//
	// }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceShrinkRequest) GetDevelop() *string {
	return s.Develop
}

func (s *CreateServiceShrinkRequest) GetLabelsShrink() *string {
	return s.LabelsShrink
}

func (s *CreateServiceShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateServiceShrinkRequest) GetBody() *string {
	return s.Body
}

func (s *CreateServiceShrinkRequest) SetDevelop(v string) *CreateServiceShrinkRequest {
	s.Develop = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetLabelsShrink(v string) *CreateServiceShrinkRequest {
	s.LabelsShrink = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetWorkspaceId(v string) *CreateServiceShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetBody(v string) *CreateServiceShrinkRequest {
	s.Body = &v
	return s
}

func (s *CreateServiceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
