/*
 * Pipeline API
 *
 * Pipeline is a feature rich application platform, built for containers on top of Kubernetes to automate the DevOps experience, continuous application development and the lifecycle of deployments. 
 *
 * API version: latest
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline
// ScanLogItem struct for ScanLogItem
type ScanLogItem struct {
	ReleaseName string `json:"releaseName,omitempty"`
	Resource string `json:"resource,omitempty"`
	Image []ScanLogItemImage `json:"image,omitempty"`
	Result []string `json:"result,omitempty"`
	Action string `json:"action,omitempty"`
}
