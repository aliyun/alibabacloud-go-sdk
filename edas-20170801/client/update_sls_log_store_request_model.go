// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSlsLogStoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateSlsLogStoreRequest
	GetAppId() *string
	SetConfigs(v string) *UpdateSlsLogStoreRequest
	GetConfigs() *string
}

type UpdateSlsLogStoreRequest struct {
	// The ID of the application. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// af58edcf-f7eb-****-****-db4e425f****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The configurations of the Logstore.
	//
	// 	- The following parameters are included in the configurations:****
	//
	//     	- **type**: the collection type. Set this parameter to file to specify the file type. Set this parameter to stdout to specify the standard output type.
	//
	//     	- **logstore**: the name of the Logstore. Make sure that the name of the Logstore is unique in the cluster. The name must comply with the following rules:
	//
	//         	- The name can contain only lowercase letters, digits, hyphens (-), and underscores (_).
	//
	//         	- The name must start and end with a lowercase letter or a digit.
	//
	//         	- The name must be 3 to 63 characters in length.
	//
	//         **
	//
	//         **Note**If you leave this parameter empty, the system automatically generates a name.
	//
	//     	- **LogDir**: If the standard output type is used, the collection path is stdout.log. If the file type is used, the collection path is the path of the collected file. Wildcards (\\*) are supported. The collection path must match the following regular expression: `^/(.+)/(.*)^/$`.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"logstore":"thisisanotherfilelog","type":"file","logDir":"/var/log/*"},{"logstore":"","type":"stdout","logDir":"stdout.log"},{"logstore":"thisisafilelog","type":"file","logDir":"/tmp/log/*"}]
	Configs *string `json:"Configs,omitempty" xml:"Configs,omitempty"`
}

func (s UpdateSlsLogStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSlsLogStoreRequest) GoString() string {
	return s.String()
}

func (s *UpdateSlsLogStoreRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateSlsLogStoreRequest) GetConfigs() *string {
	return s.Configs
}

func (s *UpdateSlsLogStoreRequest) SetAppId(v string) *UpdateSlsLogStoreRequest {
	s.AppId = &v
	return s
}

func (s *UpdateSlsLogStoreRequest) SetConfigs(v string) *UpdateSlsLogStoreRequest {
	s.Configs = &v
	return s
}

func (s *UpdateSlsLogStoreRequest) Validate() error {
	return dara.Validate(s)
}
