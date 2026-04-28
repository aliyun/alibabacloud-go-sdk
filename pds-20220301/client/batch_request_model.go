// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRequests(v []*BatchRequestRequests) *BatchRequest
	GetRequests() []*BatchRequestRequests
	SetResource(v string) *BatchRequest
	GetResource() *string
}

type BatchRequest struct {
	// The child requests.
	//
	// The number of child requests. Valid value: 1 to 100.
	//
	// This parameter is required.
	Requests []*BatchRequestRequests `json:"requests,omitempty" xml:"requests,omitempty" type:"Repeated"`
	// The type of the resource that you want to manage. Valid values:
	//
	// 	- file: a file.
	//
	// 	- drive: an individual drive or a team drive.
	//
	// 	- user: a user.
	//
	// 	- group: a group.
	//
	// 	- membership: a group member.
	//
	// 	- share_link: a share.
	//
	// 	- async_task: an asynchronous task.
	//
	// This parameter is required.
	//
	// example:
	//
	// file
	Resource *string `json:"resource,omitempty" xml:"resource,omitempty"`
}

func (s BatchRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchRequest) GoString() string {
	return s.String()
}

func (s *BatchRequest) GetRequests() []*BatchRequestRequests {
	return s.Requests
}

func (s *BatchRequest) GetResource() *string {
	return s.Resource
}

func (s *BatchRequest) SetRequests(v []*BatchRequestRequests) *BatchRequest {
	s.Requests = v
	return s
}

func (s *BatchRequest) SetResource(v string) *BatchRequest {
	s.Resource = &v
	return s
}

func (s *BatchRequest) Validate() error {
	if s.Requests != nil {
		for _, item := range s.Requests {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchRequestRequests struct {
	// The request parameters of a child request. The parameter value must be a JSON string. For more information, see the topic of the corresponding child request.
	//
	// Before you specify the request body, you must specify a header by using Content-Type. Content-Type can only be set to application/json.
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
	// The header of a child request, which indicates the type of the data specified in the request body.
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	// The ID of the child request. The ID is used to associate a child request with a response. The ID of a child request must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 93433894994ad2e1
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The method of a child request. Valid values:
	//
	// 	- POST
	//
	// 	- GET
	//
	// 	- PUT
	//
	// 	- DELETE
	//
	// 	- HEAD
	//
	// This parameter is required.
	//
	// example:
	//
	// POST
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// The API path of a child request. Valid values:
	//
	// 	- /file/get: queries the information about a file.
	//
	// 	- /file/update: modifies the information about a file.
	//
	// 	- /file/search: searches for a file.
	//
	// 	- /file/copy: copies a file or folder.
	//
	// 	- /file/move: moves a file or folder.
	//
	// 	- /file/delete: deletes a file or folder.
	//
	// 	- /file/get_download_url: queries the download URL of a file.
	//
	// 	- /file/get_share_link_download_url: queries the download URL of a file in a share.
	//
	// 	- /recyclebin/trash: moves a file or folder to the recycle bin.
	//
	// 	- /recyclebin/restore: restores a file or folder.
	//
	// 	- /file/put_usertags: adds tags to a user.
	//
	// 	- /file/delete_usertags: removes tags from a user.
	//
	// 	- /drive/get: queries the information about a drive.
	//
	// 	- /user/get: queries the information about a user.
	//
	// 	- /group/get: queries the information about a group.
	//
	// 	- /share_link/create: creates a share.
	//
	// 	- /share_link/update: modifies a share.
	//
	// 	- /share_link/cancel: cancels a share.
	//
	// 	- /share_link/list: queries shares.
	//
	// 	- /share_link/get: queries the information about a share.
	//
	// 	- /share_link/get_share_token: queries an access token of a share.
	//
	// 	- /async_task/get: queries the information about an asynchronous task.
	//
	// This parameter is required.
	//
	// example:
	//
	// /file/get
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s BatchRequestRequests) String() string {
	return dara.Prettify(s)
}

func (s BatchRequestRequests) GoString() string {
	return s.String()
}

func (s *BatchRequestRequests) GetBody() map[string]interface{} {
	return s.Body
}

func (s *BatchRequestRequests) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchRequestRequests) GetId() *string {
	return s.Id
}

func (s *BatchRequestRequests) GetMethod() *string {
	return s.Method
}

func (s *BatchRequestRequests) GetUrl() *string {
	return s.Url
}

func (s *BatchRequestRequests) SetBody(v map[string]interface{}) *BatchRequestRequests {
	s.Body = v
	return s
}

func (s *BatchRequestRequests) SetHeaders(v map[string]*string) *BatchRequestRequests {
	s.Headers = v
	return s
}

func (s *BatchRequestRequests) SetId(v string) *BatchRequestRequests {
	s.Id = &v
	return s
}

func (s *BatchRequestRequests) SetMethod(v string) *BatchRequestRequests {
	s.Method = &v
	return s
}

func (s *BatchRequestRequests) SetUrl(v string) *BatchRequestRequests {
	s.Url = &v
	return s
}

func (s *BatchRequestRequests) Validate() error {
	return dara.Validate(s)
}
