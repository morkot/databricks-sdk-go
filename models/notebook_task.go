/*
 * Databricks
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package models

type NotebookTask struct {
	NotebookPath string `json:"notebook_path"`

	BaseParameters []map[string]string `json:"base_parameters,omitempty"`
}
