// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDevelop(v string) *CreateServiceRequest
	GetDevelop() *string
	SetLabels(v map[string]*string) *CreateServiceRequest
	GetLabels() map[string]*string
	SetWorkspaceId(v string) *CreateServiceRequest
	GetWorkspaceId() *string
	SetBody(v string) *CreateServiceRequest
	GetBody() *string
}

type CreateServiceRequest struct {
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
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
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

func (s CreateServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceRequest) GetDevelop() *string {
	return s.Develop
}

func (s *CreateServiceRequest) GetLabels() map[string]*string {
	return s.Labels
}

func (s *CreateServiceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateServiceRequest) GetBody() *string {
	return s.Body
}

func (s *CreateServiceRequest) SetDevelop(v string) *CreateServiceRequest {
	s.Develop = &v
	return s
}

func (s *CreateServiceRequest) SetLabels(v map[string]*string) *CreateServiceRequest {
	s.Labels = v
	return s
}

func (s *CreateServiceRequest) SetWorkspaceId(v string) *CreateServiceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateServiceRequest) SetBody(v string) *CreateServiceRequest {
	s.Body = &v
	return s
}

func (s *CreateServiceRequest) Validate() error {
	return dara.Validate(s)
}
